package tasks

import "completerr/scheduler"

var sched = scheduler.GetScheduler()

func StartScheduler() {

	sched.Start()
	ScheduleImport()
	ScheduleSearch()
	ScheduleSonarrImport()
	ScheduleTvSearch()
}
func RestartScheduler() {
	logger.Info("Restarting Scheduler")
	sched.Clear()
	sched.Stop()
	StartScheduler()
}
