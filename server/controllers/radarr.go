package controllers

import (
	"completerr/scheduler"
	"completerr/services"
	"encoding/json"
	"net/http"
)

type MsgResp struct {
	Msg string `json:"msg"`
}

func RadarrLibraryImport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	go services.ImportRadarrMovies()
	err := json.NewEncoder(w).Encode(MsgResp{Msg: "Starting Import"})
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
func RadarrMissingSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	go scheduler.SearchMissingRadarrMovies()
	err := json.NewEncoder(w).Encode(MsgResp{Msg: "Starting Search"})
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
