package world

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/kassybas/reeu/models/country"
)

type World struct {
	Countries    []country.Country
	beginMonthWG sync.WaitGroup
}

// Init creates the world
func Init() *World {
	w := new(World)
	paths, _ := filepath.Glob("data/*/country.yaml")
	fmt.Println(paths)
	w.Countries = make([]country.Country, len(paths))
	for i, path := range paths {
		w.Countries[i] = country.NewCountry(path)
	}
	return w
}

// GetStats gets some info from a country with a givne id
func (w *World) GetStats(id int) string {
	return w.Countries[id].GetStat()
}

func (w *World) GetAllStats() string {
	s := ""
	for _, c := range w.Countries {
		s += "\n"
		s += c.GetStat()
	}
	return s
}

// Starts the "beginning of the month" update
func (w *World) StartBeginMonthUpdate() {
	for i := 0; i < len(w.Countries); i++ {
		w.beginMonthWG.Add(1)
		y := i
		go func() {
			defer w.beginMonthWG.Done()
			w.Countries[y].CollectMonthly()
		}()
	}
}

// Finishes the "beginning of the month" update
func (w *World) FinishBeginMonthpdate() {
	w.beginMonthWG.Wait()
}
