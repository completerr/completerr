package tasks

import (
	"completerr/db"
	"completerr/model"
	"completerr/scheduler"
	"completerr/services"
	"github.com/reugn/go-quartz/quartz"
	"github.com/spf13/viper"
	"time"
)

func ScheduleTvSearch() {
	logger.Info("Scheduling Sonarr Search")
	interval := viper.GetInt("sonarr.search.interval")
	scheduler.ScheduleJob(SearchMissingTvItemJob{}, quartz.NewSimpleTrigger(time.Minute*time.Duration(interval)))
}

func SearchMissingSonarrEpisode() {
	count := viper.GetInt("sonarr.search.count")
	items := db.GetRandomSearchTvItem(count)
	logger.Info(items)
	services.SearchSonarrEpisode(items)
}

func (j SearchMissingTvItemJob) Key() int {
	return model.SearchMissingSonarr
}
func (j SearchMissingTvItemJob) Execute() {
	logger.Info("Running scheduled Sonarr search")
	SearchMissingSonarrEpisode()
}
func (j SearchMissingTvItemJob) Description() string {
	return "Search missing movie from Sonarr"
}
func (j SearchMissingTvItemJob) Name() string {
	return "Sonarr Search"
}

type SearchMissingTvItemJob struct{}
