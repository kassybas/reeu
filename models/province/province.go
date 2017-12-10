package province

type Province struct {
	Name             string
	BaseTax          float32
	LocalTaxModifier float32
	LocalTaxIncome   float32
	LocalAutonomy    float32
}

// getValue of entity should be generalized TODO
// Monthly value
// https://eu4.paradoxwikis.com/Tax
func (p *Province) getTaxIncome() float32 {
	return (p.BaseTax + p.LocalTaxIncome) / 12 * (1 - p.LocalAutonomy)
}
