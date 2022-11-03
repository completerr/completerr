package db

import (
	"completerr/model"
	"completerr/utils"
	"github.com/morkid/paginate"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"time"
)

var logger = utils.GetLogger()

var db = getDb()

func InitDB() {

	logger.Info("Running Migrations")
	// Migrate the schema
	db.AutoMigrate(&model.RadarrItem{})
	db.AutoMigrate(&model.SonarrItem{})
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.SearchRecord{})

}

func getDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("completerr.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func AddItems(items []model.RadarrItem) []model.RadarrItem {

	logger.Debug("Adding RadarrItem", items)
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "radarr_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "title", "available", "released", "tmdb_id"}),
	}).CreateInBatches(&items, 100)
	if result.Error != nil {
		logger.Error(result.Error)
	}
	return items
}
func MarkAllTVItemsAsDeleted() {
	db.Delete(&model.SonarrItem{})
}
func AddTVItems(items []model.SonarrItem) []model.SonarrItem {
	logger.Debug("Adding SonarrItem", items)
	chunkSize := 100
	for _, chunk := range ChunkTV(items, chunkSize) {

		result := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "sonarr_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "title", "available", "sonarr_series_id", "series_title", "season", "season_episode_number", "absolute_episode_number"}),
		}).Create(&chunk)
		//}).CreateInBatches(&items, 100)
		if result.Error != nil {
			logger.Error(result.Error)
		}
	}
	return items
}
func ChunkTV(items []model.SonarrItem, chunkSize int) [][]model.SonarrItem {
	var chunks [][]model.SonarrItem
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}
func RemoveItem(tmdbId int64) {
	logger.Debug("Removing Id if exists")
	result := db.Where(model.RadarrItem{TMDBId: tmdbId}).Delete(&model.RadarrItem{})
	if result.Error != nil {
		logger.Error(result.Error)
	}
}
func MarkTvItemAsSearched(item *model.SonarrItem, task model.Task) {
	item.LastSearched = time.Now()
	item.SearchCount++
	db.Save(item)
	LogTvItemSearchRecord(task, *item)
}
func MarkItemAsSearched(item *model.RadarrItem, task model.Task) {
	item.LastSearched = time.Now()
	item.SearchCount++
	db.Save(item)
	LogMovieItemSearchRecord(task, *item)
}

func GetRandomSearchTvItem(count int) []model.SonarrItem {
	items := []model.SonarrItem{}
	lookbackDays, err := strconv.Atoi(viper.GetString("sonarr.search.backoff_days"))
	if err != nil {
		logger.Error(err)
	}

	lastSearchedCutoff := time.Now().Add(time.Hour * time.Duration(lookbackDays*24))
	result := db.Table("sonarr_items").Where(model.SonarrItem{Available: false}).Where("last_searched <= ?", lastSearchedCutoff).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "RANDOM()", WithoutParentheses: true},
	}).Limit(count).Scan(&items)
	if result.Error != nil {
		logger.Error(result.Error)
	}
	return items
}
func GetRandomSearchItem(count int) []model.RadarrItem {
	items := []model.RadarrItem{}
	lookbackDays, err := strconv.Atoi(viper.GetString("radarr.search.backoff_days"))
	if err != nil {
		logger.Error(err)
	}

	lastSearchedCutoff := time.Now().Add(time.Hour * time.Duration(lookbackDays*24))
	result := db.Table("radarr_items").Where(model.RadarrItem{Available: false, Released: true}).Where("last_searched <= ?", lastSearchedCutoff).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "RANDOM()", WithoutParentheses: true},
	}).Limit(count).Scan(&items)
	if result.Error != nil {
		logger.Error(result.Error)
	}
	return items
}
func LogTaskStart(jobType model.CompleterrJob) model.Task {
	task := model.Task{
		Type:     jobType.Name,
		Status:   "started",
		Started:  time.Now(),
		Finished: time.Time{},
	}
	db.Create(&task)
	return task
}
func LogTaskFinish(task model.Task) model.Task {
	task.Finished = time.Now()
	task.Status = "finished"
	db.Save(&task)
	return task
}
func LogMovieItemSearchRecord(task model.Task, item model.RadarrItem) model.SearchRecord {
	searchRecord := model.SearchRecord{RadarrItem: item, Task: task}
	db.Create(&searchRecord)
	return searchRecord
}
func LogTvItemSearchRecord(task model.Task, item model.SonarrItem) model.SearchRecord {
	searchRecord := model.SearchRecord{SonarrItem: item, Task: task}
	db.Create(&searchRecord)
	return searchRecord
}

func GetTaskHistory(r *http.Request) paginate.Page {
	var tasks = []model.Task{}
	model := db.Model(&model.Task{}).Order("created_at desc")
	pg := paginate.New()
	page := pg.With(model).Request(r).Response(&tasks)

	return page
}
func GetMostRecentTaskRun(job model.CompleterrJob) model.Task {
	var task = model.Task{}
	db.Model(model.Task{}).Order("created_at desc").Where("type = ?", job.Name).First(&task)
	return task
}

func GetSearchHistory(r *http.Request, includeMovie bool, includeTv bool) paginate.Page {
	var searchRecords = []model.SearchRecord{}
	model := db.Model(&model.SearchRecord{}).Order("created_at desc")
	if includeMovie {
		model.Where("radarr_item_id > 0").Preload("RadarrItem")
	} else if includeTv {
		model.Where("sonarr_item_id > 0").Preload("SonarrItem")
	}
	pg := paginate.New()
	page := pg.With(model).Request(r).Response(&searchRecords)

	return page
}
