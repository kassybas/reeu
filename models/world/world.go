package world

import (
	"sync"

	"github.com/kassybas/reeu/models/country"
)

type World struct {
	Countries []country.Country
	ym1WG     sync.WaitGroup
}

func Init() *World {
	w := new(World)
	w.Countries = make([]country.Country, 1)
	w.Countries[0] = country.NewCountry("data/sweeden.yaml")
	return w
}

func (w *World) GetStats(id int) string {
	return w.Countries[id].GetStat()
}

func (w *World) StartYM1Update() {
	for i := 0; i < len(w.Countries); i++ {
		w.ym1WG.Add(1)
		y := i
		go func() {
			defer w.ym1WG.Done()
			w.Countries[y].CollectTaxes()
			w.Countries[y].CollectProduction()
		}()
	}
}

func (w *World) FinishYM1Update() {
	w.ym1WG.Wait()
}
