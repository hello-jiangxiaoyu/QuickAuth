package cronjob

import (
	"QuickAuth/internal/global"
	"QuickAuth/pkg/utils"
	"time"
)

func StartDailyInspectionTask(hour int, minute int, fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil && global.Config != nil {
				global.Log.Error(utils.GetPanicStackInfo("DailyTimer", err, 3, global.Config.Log.IsFullStack))
			}
		}()
		for {
			now := time.Now()
			next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
			if next.Sub(now) < 1*time.Minute {
				next = next.Add(24 * time.Hour)
			}
			t := time.NewTimer(next.Sub(now))
			<-t.C

			fn()
		}
	}()
}
