package cronjob

import (
	"QuickAuth/internal/model"
	"QuickAuth/pkg/global"
	"go.uber.org/zap"
	"time"
)

func StartClearExpiredCodeInspection() {
	StartDailyInspectionTask(4, 0, func() {
		earliest := time.Now().Add(-2 * time.Minute)
		if err := global.DB.Where("created_at < ?", earliest).Delete(model.Code{}).Error; err != nil {
			global.Log.Error("clear expired code err: ", zap.Error(err))
		}
	})
}
