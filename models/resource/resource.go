package resource

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Resource is a generic representation of everything that is collectable (money, points, production, gold, etc)
// Parts are to be summed: sum(part[i].collect())
// Modifiers are multiplied with each other and with the sum
// Flat is the value of the bottom resource. It should be non-zero only if other resources are done
type Resource struct {
	Name          string `yaml:"Name"`
	Parts         []Resource
	PartsPath     []string   `yaml:"PartsPath"`
	Modifiers     []Modifier `yaml:"Modifiers"`
	ModifiersPath []string   `yaml:"ModifierPath"`
	Flat          float64    `yaml:"Flat"`
}

type Modifier struct {
	Name   string  `yaml:"Name"`
	Amount float64 `yaml:"Amount"`
}

// NewResource factory method to create a New Resource
func NewResource(name string, parts []Resource, modifiers []Modifier, flat float64) Resource {
	r := new(Resource)
	r.Name = name
	r.Parts = parts
	r.Modifiers = modifiers
	r.Flat = flat
	return *r
}

// LoadResource from a file
func LoadResource(path string) Resource {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := new(Resource)
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	for _, m := range c.ModifiersPath {
		c.Modifiers = append(c.Modifiers, loadModifier(m))
	}

	c.Parts = make([]Resource, len(c.PartsPath))
	for i, p := range c.PartsPath {
		c.Parts[i] = LoadResource(p)
	}
	return *c
}

func loadModifier(path string) Modifier {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	m := new(Modifier)
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return *m
}

// CollectMonthly gives the monthly amount of collectables
func (r *Resource) CollectMonthly() float64 {
	return r.Collect() / 12
}

// Collect resource periodically
func (r *Resource) Collect() float64 {
	sum := r.Flat
	for _, p := range r.Parts {
		sum += p.Collect()
	}
	for _, m := range r.Modifiers {
		sum *= m.Amount
	}
	return sum
}
