package cronjob

func StartDeleteDefunctUserPoolInspection() {
	StartDailyInspectionTask(4, 30, func() {

	})
}
