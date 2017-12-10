package country

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kassybas/reeu/models/province"

	yaml "gopkg.in/yaml.v2"
)

type Country struct {
	Name                 string   `yaml:"Name"`
	ProvinceList         []string `yaml:"ProvinceList"`
	Provinces            []province.Province
	Money                float64 `yaml:"Money"`
	TaxIncomeEfficency   float64 `yaml:"TaxIncomeEfficiency"`
	ProductionEfficiency float64 `yaml:"ProductionEfficiency"`
}

func loadCountry(path string) *Country {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := new(Country)
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
func NewCountry(path string) Country {
	c := loadCountry(path)
	c.Provinces = make([]province.Province, len(c.ProvinceList))
	for i, path := range c.ProvinceList {
		c.Provinces[i] = province.NewProvince(path)
	}
	return *c
}

func (c *Country) CollectTaxes() {
	for _, p := range c.Provinces {
		c.Money += p.GetTaxIncome() * c.TaxIncomeEfficency
	}
}

func (c *Country) CollectProduction() {
	for _, p := range c.Provinces {
		c.Money += p.GetProductionIncome() //* c.ProductionEfficiency
	}
}

func (c *Country) GetStat() string {
	return "---\n" + "Name: " + c.Name +
		fmt.Sprintf("\nMoney: %.2f", c.Money) +
		"\n---\n"
}

// TODO: modifiers for all, buildings, technology
