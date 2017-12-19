package country

import (
	"log"
	"path/filepath"

	"github.com/kassybas/reeu/models/common"

	"github.com/kassybas/reeu/models/resource"

	yaml "gopkg.in/yaml.v2"
)

type Country struct {
	Name          string `yaml:"Name"`
	basePath      string
	ResourcePaths []string `yaml:"ResourcePaths"`
	Resources     []resource.Resource
}

func loadCountry(path string) Country {
	yamlFile := common.LoadFile(path)
	c := new(Country)
	err := yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	c.basePath = filepath.Dir(path) + "/"
	return *c
}

func NewCountry(path string) Country {
	c := loadCountry(path)
	c.Resources = make([]resource.Resource, len(c.ResourcePaths))
	for i, p := range c.ResourcePaths {
		c.Resources[i] = resource.LoadResource(c.basePath, p)
	}
	return c
}

func (c *Country) CollectMonthly() {
	for i := range c.Resources {
		c.Resources[i].CollectMonthly()
	}
}

func (c *Country) GetStat() string {
	s := "---\n"
	s += c.Name + "\n"
	for _, r := range c.Resources {
		s += r.GetStat() + "\n"
	}
	return s
}
