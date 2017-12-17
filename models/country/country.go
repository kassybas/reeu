package country

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/kassybas/reeu/models/resource"

	yaml "gopkg.in/yaml.v2"
)

type Country struct {
	Name               string  `yaml:"Name"`
	Money              float64 `yaml:"Money"`
	Path               string
	TaxResource        resource.Resource
	ProductionResource resource.Resource
}

func loadCountry(path string) Country {
	c := new(Country)
	c.Path = path
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return *c
}
func NewCountry(path string) Country {
	c := loadCountry(path)
	c.TaxResource = resource.LoadResource(filepath.Dir(c.Path)+"/", "tax/main.yaml")
	c.ProductionResource = resource.LoadResource(filepath.Dir(c.Path)+"/", "production/main.yaml")
	return c
}

func (c *Country) CollectMonthly() {
	c.Money += c.ProductionResource.CollectMonthly()
	c.Money += c.TaxResource.CollectMonthly()
}

func (c *Country) GetStat() string {
	return "---\n" + "Name: " + c.Name +
		fmt.Sprintf("\nMoney: %.2f", c.Money) +
		"\n---\n"
}
