package scheduler

import (
	"completerr/utils"
	"github.com/reugn/go-quartz/quartz"
)

var logger = utils.GetLogger()

var sched = quartz.NewStdScheduler()

func GetScheduler() quartz.StdScheduler {
	return *sched
}

func ScheduleJob(job quartz.Job, trigger quartz.Trigger) error {
	return sched.ScheduleJob(job, trigger)
}
