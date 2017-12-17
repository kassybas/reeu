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
	Name string `yaml:"Name"`
	Path string
	//TODO: generalize resource and init values as in source
	Money      resource.Resource
	StartMoney float64 `yaml:"Money"`

	Manpower      resource.Resource
	StartManpower float64 `yaml:"Manpower"`
	MaxManpower   float64 `yaml:"MaxManpower"`
}

func loadCountry(path string) Country {
	c := new(Country)
	c.Path = filepath.Dir(path) + "/"
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

	sources := []string{"tax/main.yaml", "production/main.yaml"}
	c.Money = resource.NewResource("Money", c.Path, sources, c.StartMoney, 0.0)

	sources = []string{"manpower/main.yaml"}
	c.Manpower = resource.NewResource("Manpower", c.Path, sources, c.StartManpower, c.MaxManpower)

	return c
}

func (c *Country) CollectMonthly() {
	c.Money.CollectMonthly()
	c.Manpower.CollectMonthly()
}

func (c *Country) GetStat() string {
	return "---\n" + "Name: " + c.Name +
		fmt.Sprintf("\nMoney: %.2f", c.Money.Amount) +
		fmt.Sprintf("\nManpower: %.2f", c.Manpower.Amount) +
		"\n---\n"
}
