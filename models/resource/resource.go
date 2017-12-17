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
	Group  string  `yaml:"Group"`
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
func LoadResource(basePath, path string) Resource {
	yamlFile, err := ioutil.ReadFile(basePath + path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := new(Resource)
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	for _, m := range c.ModifiersPath {
		c.Modifiers = append(c.Modifiers, loadModifier(basePath, m))
	}
	c.Parts = make([]Resource, len(c.PartsPath))
	for i, p := range c.PartsPath {
		c.Parts[i] = LoadResource(basePath, p)
	}
	return *c
}

func loadModifier(basePath, path string) Modifier {
	yamlFile, err := ioutil.ReadFile(basePath + path)
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

func (r *Resource) getModifierProduct() float64 {
	mod := 1.0
	groups := make(map[string]float64)
	for _, m := range r.Modifiers {
		if m.Group == "" {
			mod *= m.Amount
		} else {
			groups[m.Group] += m.Amount
		}
	}
	// The product of the modifier groups
	for _, value := range groups {
		// When "Group" key is present, the values are not in _absolute_ percentage (eg.: 1.15 for 115%) but rather _relative_ increase (0.15 for +15%)
		// This has to be compensated at the end for the correct result (so +15% is not decreasing the actual value)
		value += 1.0
		mod *= value
	}
	return mod
}

// CollectMonthly gives the monthly amount of collectables
func (r *Resource) CollectMonthly() float64 {
	return r.Collect() / 12
}

// Collect resource periodically
func (r *Resource) Collect() float64 {
	sum := r.Flat
	// Get the parts
	for _, p := range r.Parts {
		sum += p.Collect()
	}
	mod := r.getModifierProduct()

	return sum * mod
}
