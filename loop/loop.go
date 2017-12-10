package loop

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/kassybas/reeu/models/world"
)

type loopConfig struct {
	startDate time.Time
	endDate   time.Time
}

func setupConfig() loopConfig {
	return loopConfig{
		startDate: time.Date(2005, 1, 17, 0, 0, 0, 0, time.UTC),
		endDate:   time.Date(2006, 12, 20, 0, 0, 0, 0, time.UTC),
	}
}

func triggerYM1(w *world.World) {
	w.StartYM1Update()
}

func triggerYM10(w *world.World) {
	w.FinishYM1Update()
}

func StartLoop() {
	reader := bufio.NewReader(os.Stdin)
	lc := setupConfig()
	w := world.Init()
	fmt.Printf(w.GetStats(0))
	reader.ReadString('\n')

	for cur := lc.startDate; cur.Before(lc.endDate); cur = cur.AddDate(0, 0, 1) {
		if cur.Day() == 1 {
			fmt.Printf(" - ")
			fmt.Println(cur.Date())
			triggerYM1(w)
		}
		if cur.Day() == 10 {
			triggerYM10(w)
		}

		if cur.Month() == 1 && cur.Day() == 1 {
			fmt.Println("===== [[ Yearly Pass ]] =====")
			fmt.Printf(w.GetStats(0))
			reader.ReadString('\n')
		}
	}
	fmt.Printf(w.GetStats(0))
}
