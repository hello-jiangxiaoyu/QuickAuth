package cronjob

import (
	"QuickAuth/internal/global"
	"QuickAuth/pkg/model"
	"go.uber.org/zap"
	"time"
)

func ClearExpiredCode() {
	go func() {
		for {
			now := time.Now()
			next := time.Date(now.Year(), now.Month(), now.Day(), 4, 0, 0, 0, now.Location())
			if next.Sub(now) < 1*time.Hour {
				next = next.Add(24 * time.Hour)
			}
			t := time.NewTimer(next.Sub(now))
			<-t.C

			earliest := time.Now().Add(-2 * time.Minute)
			if err := global.DB.Where("created_at < ?", earliest).Delete(model.Code{}).Error; err != nil {
				global.Log.Error("clear expired code err: ", zap.Error(err))
			}
		}
	}()
}
