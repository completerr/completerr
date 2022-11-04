package controllers

import (
	"completerr/db"
	"completerr/model"
	"completerr/services"
	"completerr/utils"
	"encoding/json"
	"net/http"
)

var logger = utils.GetLogger()

type TaskInfoResponse map[string]services.TaskInfo

func TaskInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	infoMap := make(map[string]services.TaskInfo)
	for _, job := range model.CompleterrJobs {
		info, err := services.GetTaskInfo(job)
		if err != nil {
			logger.Error(err)
			continue
		}
		infoMap[job.ResponseKey] = info
	}
	json.NewEncoder(w).Encode(infoMap)
}

func TaskHistory(w http.ResponseWriter, r *http.Request) {
	tasks := db.GetTaskHistory(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
