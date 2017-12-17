package resource

import "github.com/kassybas/reeu/models/source"

type Resource struct {
	Name      string
	Sources   []source.Source
	Amount    float64
	MaxAmount float64
}

func NewResource(name string, basePath string, sourcePaths []string, amount, maxAmount float64) Resource {
	r := new(Resource)
	r.Name = name
	r.Amount = amount
	r.MaxAmount = maxAmount
	r.Sources = make([]source.Source, len(sourcePaths))
	for i, s := range sourcePaths {
		r.Sources[i] = source.LoadSource(basePath, s)
	}
	return *r
}

func (r *Resource) CollectMonthly() {
	for _, s := range r.Sources {
		r.Amount += s.CollectMonthly()
		if r.MaxAmount != 0 && r.MaxAmount < r.Amount {
			r.Amount = r.MaxAmount
			break
		}
	}
}
