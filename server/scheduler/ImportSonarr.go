package scheduler

import (
	"completerr/services"
	"github.com/reugn/go-quartz/quartz"
	"github.com/spf13/viper"
	"time"
)

func ScheduleSonarrImport() {
	logger.Info("Scheduling Import")
	interval := viper.GetInt("sonarr.library_sync_interval")
	sched.ScheduleJob(ImportSonarrJob{}, quartz.NewSimpleTrigger(time.Minute*time.Duration(interval)))
}

func ImportSonarrMovies() {
	services.ImportSonarrEpisodes()
}

func (j ImportSonarrJob) Key() int {
	return ImportSonarr
}
func (j ImportSonarrJob) Execute() {
	logger.Info("Running scheduled Sonarr import")
	ImportSonarrMovies()
}
func (j ImportSonarrJob) Description() string {
	return "Sync missing movies from Sonarr"
}

func (j ImportSonarrJob) Name() string {
	return "Sonarr Import"
}

type ImportSonarrJob struct{}
