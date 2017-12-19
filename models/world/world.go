package world

import (
	"path/filepath"
	"sync"
	"time"

	"github.com/kassybas/reeu/models/country"
)

type World struct {
	StartDate    time.Time
	EndDate      time.Time
	CurDate      time.Time
	Countries    []country.Country
	beginMonthWG sync.WaitGroup
}

// Init creates the world
func Init() *World {
	w := new(World)
	w.StartDate = time.Date(1444, 1, 1, 0, 0, 0, 0, time.UTC)
	w.EndDate = time.Date(2000, 12, 20, 0, 0, 0, 0, time.UTC)
	paths, _ := filepath.Glob("data/*/country.yaml")
	w.Countries = make([]country.Country, len(paths))
	for i, path := range paths {
		w.Countries[i] = country.NewCountry(path)
	}
	return w
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
