package resource

import (
	"fmt"
	"log"

	"github.com/kassybas/reeu/models/common"

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
	Keep          float64    `yaml:"Keep"`
	Max           float64    `yaml:"Max"`
	MaxPath       string     `yaml:"MaxPath"`
	MaxResource   *Resource
	Stored        float64 `yaml:"Stored"`
}

type Modifier struct {
	Name   string  `yaml:"Name"`
	Amount float64 `yaml:"Amount"`
	Group  string  `yaml:"Group"`
}

// LoadSource from a file
func LoadResource(basePath, path string) Resource {
	yamlFile := common.LoadFile(basePath + path)
	c := new(Resource)
	err := yaml.Unmarshal(yamlFile, c)
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
	if c.Max != 0 {
		mr := LoadResource(basePath, c.MaxPath)
		c.MaxResource = &mr
		c.Max = c.MaxResource.Collect()
	}
	return *c
}

func loadModifier(basePath, path string) Modifier {
	yamlFile := common.LoadFile(basePath + path)
	m := new(Modifier)
	err := yaml.Unmarshal(yamlFile, m)
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
			// Groups percentages should be added relative: 1.15(+15%) + 1.2(+20%) = 1.35(+35%) | NOT: 235% )
			groups[m.Group] += (m.Amount - 1)
		}
	}
	// Get the product of the modifier-groups
	for _, value := range groups {
		value += 1
		mod *= value
	}
	return mod
}

// CollectMonthly gives the monthly amount of collectables
func (r *Resource) CollectMonthly() float64 {
	inc := r.Collect() / 12
	r.Stored += inc * r.Keep
	if r.Max != 0 {
		r.Max = r.MaxResource.Collect()
		if r.Max < r.Stored {
			r.Stored = r.Max
		}
	}
	return inc * (1 - r.Keep)
}

// Collect resource periodically
func (r *Resource) Collect() float64 {
	sum := r.Flat
	for _, p := range r.Parts {
		sum += p.Collect()
	}
	mod := r.getModifierProduct()
	return sum * mod
}

func (r *Resource) GetStat() string {
	return r.Name + ": " + fmt.Sprint(r.Stored)
}
