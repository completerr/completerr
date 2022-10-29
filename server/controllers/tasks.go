package controllers

import (
	"completerr/db"
	"completerr/scheduler"
	"completerr/utils"
	"encoding/json"
	"github.com/reugn/go-quartz/quartz"
	"net/http"
)

var logger = utils.GetLogger()

type TaskInfoResponse struct {
	RadarrImport quartz.ScheduledJob `json:"radarr_import"`
}

func TaskInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sched := scheduler.GetScheduler()
	rImport, err := sched.GetScheduledJob(scheduler.ImportRadarr)
	if err != nil {
		logger.Error(err)
	}
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
	err = json.NewEncoder(w).Encode(TaskInfoResponse{RadarrImport: *rImport})
}
func TaskHistory(w http.ResponseWriter, r *http.Request) {
	tasks := db.GetTaskHistory(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
