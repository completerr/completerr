package db

import (
	"completerr/model"
	"completerr/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

var logger = utils.GetLogger()

var db = getDb()

func InitDB() {

	logger.Info("Running Migrations")
	// Migrate the schema
	db.AutoMigrate(&model.Item{})
	db.AutoMigrate(&model.TvItem{})
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.SearchRecord{})

	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}

func getDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("completerr.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func AddItems(items []model.Item) []model.Item {

	logger.Debug("Adding Item", items)
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
	db.Delete(&model.TvItem{})
}
func AddTVItems(items []model.TvItem) []model.TvItem {
	logger.Debug("Adding Item", items)
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
func ChunkTV(items []model.TvItem, chunkSize int) [][]model.TvItem {
	var chunks [][]model.TvItem
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}
func RemoveItem(tmdbId int64) {
	logger.Debug("Removing Id if exists")
	result := db.Where(model.Item{TMDBId: tmdbId}).Delete(&model.Item{})
	if result.Error != nil {
		logger.Error(result.Error)
	}
}
func MarkTvItemAsSearched(item *model.TvItem, task model.Task) {
	item.LastSearched = time.Now()
	item.SearchCount++
	db.Save(item)
	LogTvItemSearchRecord(task, *item)
}
func MarkItemAsSearched(item *model.Item, task model.Task) {
	item.LastSearched = time.Now()
	item.SearchCount++
	db.Save(item)
	LogMovieItemSearchRecord(task, *item)
}

func GetRandomSearchTvItem(count int) []model.TvItem {
	items := []model.TvItem{}
	lookbackDays, err := strconv.Atoi(viper.GetString("sonarr.search.backoff_days"))
	if err != nil {
		logger.Error(err)
	}

	lastSearchedCutoff := time.Now().Add(time.Hour * time.Duration(lookbackDays*24))
	result := db.Table("tv_items").Where(model.TvItem{Available: false}).Where("last_searched <= ?", lastSearchedCutoff).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "RANDOM()", WithoutParentheses: true},
	}).Limit(count).Scan(&items)
	if result.Error != nil {
		logger.Error(result.Error)
	}
	return items
}
func GetRandomSearchItem(count int) []model.Item {
	items := []model.Item{}
	lookbackDays, err := strconv.Atoi(viper.GetString("radarr.search.backoff_days"))
	if err != nil {
		logger.Error(err)
	}

	lastSearchedCutoff := time.Now().Add(time.Hour * time.Duration(lookbackDays*24))
	result := db.Table("items").Where(model.Item{Available: false, Released: true}).Where("last_searched <= ?", lastSearchedCutoff).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "RANDOM()", WithoutParentheses: true},
	}).Limit(count).Scan(&items)
	if result.Error != nil {
		logger.Error(result.Error)
	}
	return items
}
func LogTaskStart(jobType string) model.Task {
	task := model.Task{
		Type:     jobType,
		Status:   "started",
		Started:  time.Now(),
		Finished: time.Time{},
	}
	db.Create(&task)
	return task
}
func LogTaskFinish(task model.Task) model.Task {
	task.Finished = time.Now()
	db.Save(&task)
	return task
}
func LogMovieItemSearchRecord(task model.Task, item model.Item) model.SearchRecord {
	searchRecord := model.SearchRecord{ItemID: item.ID, TaskID: task.ID}
	db.Create(&searchRecord)
	return searchRecord
}
func LogTvItemSearchRecord(task model.Task, item model.TvItem) model.SearchRecord {
	searchRecord := model.SearchRecord{ItemID: item.ID, TvItemID: task.ID}
	db.Create(&searchRecord)
	return searchRecord
}
