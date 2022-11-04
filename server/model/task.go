package model

import (
	"completerr/utils"
)

var logger = utils.GetLogger()

const (
	ImportRadarr int = iota
	ImportSonarr
	SearchMissingRadarr
	SearchMissingSonarr
)

type CompleterrJob struct {
	SchedulerKey int
	Name         string
	ResponseKey  string
}

const RADARR_IMPORT_NAME = "Radarr Import"
const SONARR_IMPORT_NAME = "Sonarr Import"
const RADARR_SEARCH_NAME = "Radarr Search"
const SONARR_SEARCH_NAME = "Sonarr Search"

var RadarrImport = CompleterrJob{
	Name:         RADARR_IMPORT_NAME,
	SchedulerKey: ImportRadarr,
	ResponseKey:  "radarr_import",
}
var SonarrImport = CompleterrJob{
	Name:         SONARR_IMPORT_NAME,
	SchedulerKey: ImportSonarr,
	ResponseKey:  "sonarr_import",
}
var RadarrSearch = CompleterrJob{
	Name:         RADARR_SEARCH_NAME,
	SchedulerKey: SearchMissingRadarr,
	ResponseKey:  "radarr_search",
}
var SonarrSearch = CompleterrJob{
	Name:         SONARR_SEARCH_NAME,
	SchedulerKey: SearchMissingSonarr,
	ResponseKey:  "sonarr_search",
}
var CompleterrJobs = []CompleterrJob{
	RadarrImport,
	SonarrImport,
	SonarrSearch,
	RadarrSearch,
}
