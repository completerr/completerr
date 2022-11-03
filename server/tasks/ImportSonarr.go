package tasks

import (
	"completerr/model"
	"completerr/scheduler"
	"completerr/services"
	"completerr/utils"
	"github.com/reugn/go-quartz/quartz"
	"github.com/spf13/viper"
	"time"
)

var logger = utils.GetLogger()

func ScheduleSonarrImport() {
	logger.Info("Scheduling Import")
	interval := viper.GetInt("sonarr.library_sync_interval")
	scheduler.ScheduleJob(ImportSonarrJob{}, quartz.NewSimpleTrigger(time.Minute*time.Duration(interval)))
}

func ImportSonarrMovies() {
	services.ImportSonarrEpisodes()
}

func (j ImportSonarrJob) Key() int {
	return model.ImportSonarr
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
