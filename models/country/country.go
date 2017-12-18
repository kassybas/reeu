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
	Name         string `yaml:"Name"`
	basePath     string
	Money        resource.Resource
	MoneyPath    string `yaml:"Money"`
	Manpower     resource.Resource
	ManpowerPath string `yaml:"Manpower"`
}

func loadCountry(path string) Country {
	c := new(Country)
	c.basePath = filepath.Dir(path) + "/"
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
	c.Money = resource.LoadResource(c.basePath, c.MoneyPath)
	c.Manpower = resource.LoadResource(c.basePath, c.ManpowerPath)
	return c
}

func (c *Country) CollectMonthly() {
	c.Money.CollectMonthly()
	c.Manpower.CollectMonthly()
}

func (c *Country) GetStat() string {
	return "---\n" + "Name: " + c.Name +
		fmt.Sprintf("\nMoney: %.2f", c.Money.Stored) +
		fmt.Sprintf("\nManpower: %.2f", c.Manpower.Stored) +
		"\n---\n"
}
