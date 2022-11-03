package controllers

import (
	"completerr/db"
	"completerr/services"
	"completerr/tasks"
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
	go tasks.SearchMissingSonarrEpisode()
	err := json.NewEncoder(w).Encode(MsgResp{Msg: "Starting Search"})
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
func SonarrSearchHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	results := db.GetSearchHistory(r, false, true)
	err := json.NewEncoder(w).Encode(results)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}
