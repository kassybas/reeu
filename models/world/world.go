package world

import (
	"fmt"
	"sync"

	"github.com/kassybas/reeu/models/country"
)

type World struct {
	Countries []country.Country
	ym1WG     sync.WaitGroup
}

func Init() *World {
	w := new(World)
	//w.ym1WG = new(sync.WaitGroup)
	w.Countries = make([]country.Country, 1)
	w.Countries[0] = country.NewCountry()
	// w.Countries[1].Money = 12
	// w.Countries[1].Name = "Hungary"
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
		}()
	}
}

func (w *World) FinishYM1Update() {
	fmt.Println("Were waiting")
	w.ym1WG.Wait()
}
