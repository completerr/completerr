package scheduler

import (
	"completerr/utils"
	"github.com/reugn/go-quartz/quartz"
)

var logger = utils.GetLogger()

const (
	ImportRadarr int = iota
	ImportSonarr
	SearchMissingRadarr
	SearchMissingSonarr
)

var sched = quartz.NewStdScheduler()

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
func GetScheduler() quartz.StdScheduler {
	return *sched
}
