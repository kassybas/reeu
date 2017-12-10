package province

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Province struct {
	Name          string  `yaml:"Name"`
	LocalAutonomy float64 `yaml:"LocalAutonomy"`
	// TAX
	BaseTax                  float64 `yaml:"BaseTax"`
	LocalTaxIncome           float64 `yaml:"LocalTaxIncome"`
	LocalTaxIncomeEfficiency float64 `yaml:"LocalTaxIncomeEfficiency"`
	TaxModifiers             []Modifier
	// Production
	ProductionValue float64 `yaml:"ProductionValue"`
	MarketPrice     float64 `yaml:"MarketPrice"`
}

type Modifier struct {
	Name        string
	Amount      float64
	TaxModifier bool
}

func loadProvince(path string) *Province {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := new(Province)
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func NewProvince(path string) Province {
	p := loadProvince(path)
	return *p
}

func (p *Province) getTaxModifiers() float64 {
	mod := 1.0
	for _, m := range p.TaxModifiers {
		mod *= m.Amount
	}
	return mod
}

// getValue of entity should be generalized TODO
// Monthly value
// https://eu4.paradoxwikis.com/Tax
func (p *Province) GetTaxIncome() float64 {
	return (p.BaseTax + p.LocalTaxIncome) / 12 * (1 - p.LocalAutonomy) * p.LocalTaxIncomeEfficiency * p.getTaxModifiers()
}

// Monthly value
// https://eu4.paradoxwikis.com/Tax
func (p *Province) GetProductionIncome() float64 {
	goodsProduced := p.ProductionValue * 0.2    //TODO: + flatValueModifiers) *goodsProducedModifiers
	tradeValue := goodsProduced * p.MarketPrice //TODO: * eventPriceModifiers
	return tradeValue * (1 - p.LocalAutonomy)   //TODO: * goodsProducedModifiers
}
