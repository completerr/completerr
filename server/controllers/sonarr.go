package controllers

import (
	"completerr/scheduler"
	"completerr/services"
	"encoding/json"
	"net/http"
)

func SonarrLibraryImport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	go services.ImportSonarrEpisodes()
	err := json.NewEncoder(w).Encode(MsgResp{Msg: "Starting Import"})
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func SonarrMissingSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	go scheduler.SearchMissingSonarrEpisode()
	err := json.NewEncoder(w).Encode(MsgResp{Msg: "Starting Search"})
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
