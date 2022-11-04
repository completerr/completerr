package services

import (
	"completerr/db"
	"completerr/model"
	"completerr/scheduler"
	"time"
)

type TaskInfo struct {
	PrevRunAt time.Time `json:"prev_run_at"`
	NextRunAt time.Time `json:"next_run_at"`
}

func GetTaskInfo(job model.CompleterrJob) (TaskInfo, error) {

	sched := scheduler.GetScheduler()
	schedJob, err := sched.GetScheduledJob(job.SchedulerKey)
	if err != nil {
		logger.Error(err)
		return TaskInfo{}, err
	}
	next := time.Unix(0, schedJob.NextRunTime)

	prevTask := db.GetMostRecentTaskRun(job)
	prev := prevTask.Started
	return TaskInfo{NextRunAt: next, PrevRunAt: prev}, nil

}
