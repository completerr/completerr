package scheduler

import (
	"completerr/db"
	"completerr/services"
	"github.com/reugn/go-quartz/quartz"
	"github.com/spf13/viper"
	"time"
)

func ScheduleSearch() {
	logger.Info("Scheduling Radarr Search")
	interval := viper.GetInt("radarr.search.interval")
	sched.ScheduleJob(SearchMissingItemJob{}, quartz.NewSimpleTrigger(time.Minute*time.Duration(interval)))
}

func SearchMissingRadarrMovies() {
	count := viper.GetInt("radarr.search.count")
	items := db.GetRandomSearchItem(count)
	logger.Info(items)
	services.SearchRadarrMovies(items)
}

func (j SearchMissingItemJob) Key() int {
	return SearchMissingRadarr
}
func (j SearchMissingItemJob) Execute() {
	logger.Info("Running scheduled Radarr search")
	SearchMissingRadarrMovies()
}
func (j SearchMissingItemJob) Description() string {
	return "Search missing movie from Radarr"
}

func (j SearchMissingItemJob) Name() string {
	return "Radarr Search"
}

type SearchMissingItemJob struct{}
