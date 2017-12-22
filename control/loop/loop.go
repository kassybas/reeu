package loop

import (
	"fmt"
	"time"

	"github.com/kassybas/reeu/models/world"
)

type loopConfig struct {
	startDate time.Time
	endDate   time.Time
}

func StartLoop() {
	w := world.Init()
	fmt.Printf(w.GetAllStats())
	for w.CurDate = w.StartDate; w.CurDate.Before(w.EndDate); w.CurDate = w.CurDate.AddDate(0, 0, 1) {
		if w.CurDate.Month() == 1 && w.CurDate.Day() == 1 {
			fmt.Println("===== [[ Yearly Pass ]] =====")
			fmt.Println(w.CurDate.Date())
			fmt.Printf(w.GetAllStats())
		}
		if w.CurDate.Day() == 1 {
			w.StartBeginMonthUpdate()
		}
		if w.CurDate.Day() == 10 {
			w.FinishBeginMonthpdate()
		}
		time.Sleep(time.Millisecond)
	}
}
