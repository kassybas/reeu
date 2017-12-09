package loop

import "fmt"
import "time"

type loopConfig struct {
	startDate time.Time
}

func setupConfig() loopConfig {
	lc = loopConfig { startDate =time.Now() }
	return time.Now()
}
func StartLoop() {
	fmt.Println("ok")

}
