package scheduler

import (
	"completerr/services"
	"github.com/reugn/go-quartz/quartz"
	"github.com/spf13/viper"
	"time"
)

func ScheduleImport() {
	logger.Info("Scheduling Import")
	interval := viper.GetInt("radarr.library_sync_interval")
	sched.ScheduleJob(ImportRadarrJob{}, quartz.NewSimpleTrigger(time.Minute*time.Duration(interval)))
}

func ImportRadarrMovies() {
	services.ImportRadarrMovies()
}

func (j ImportRadarrJob) Key() int {
	return ImportRadarr
}
func (j ImportRadarrJob) Execute() {
	logger.Info("Running scheduled Radarr import")
	ImportRadarrMovies()
}
func (j ImportRadarrJob) Description() string {
	return "Sync missing movies from Radarr"
}
func (j ImportRadarrJob) Name() string {
	return "Radarr Import"
}

type ImportRadarrJob struct{}
