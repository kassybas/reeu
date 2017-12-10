package country

import (
	"fmt"

	"github.com/kassybas/reeu/models/province"
)

type Country struct {
	Name      string
	Provinces []province.Province
	Money     float32
}

func NewCountry() Country {
	c := new(Country)
	c.Name = "Sweeden"
	c.Provinces = make([]province.Province, 2)
	c.Provinces
	c.Money = 100
	return *c
}

func (c *Country) CollectTaxes() {
	c.Money += 2
}

func (c *Country) GetStat() string {
	return "---\n" + "Name: " + c.Name +
		fmt.Sprintf("\nMoney: %.2f", c.Money) +
		"\n---\n"
}
