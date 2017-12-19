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

func StartLoop() {
	reader := bufio.NewReader(os.Stdin)
	w := world.Init()
	fmt.Printf(w.GetAllStats())
	for w.CurDate = w.StartDate; w.CurDate.Before(w.EndDate); w.CurDate = w.CurDate.AddDate(0, 0, 1) {
		if w.CurDate.Month() == 1 && w.CurDate.Day() == 1 {
			fmt.Println("===== [[ Yearly Pass ]] =====")
			fmt.Println(w.CurDate.Date())
			fmt.Printf(w.GetAllStats())
		}
		if w.CurDate.Year()%100 == 0 {
			reader.ReadString('\n')
		}
		if w.CurDate.Day() == 1 {
			w.StartBeginMonthUpdate()
		}
		if w.CurDate.Day() == 10 {
			w.FinishBeginMonthpdate()
		}
		time.Sleep(5 * time.Millisecond)
	}
}
