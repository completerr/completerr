package services

import (
	"completerr/db"
	"completerr/model"
	"completerr/utils"
	"context"
	"fmt"
	"github.com/spf13/viper"
	//"golift.io/starr"
	//"golift.io/starr/radarr"
	"github.com/craigjmidwinter/starr"
	"github.com/craigjmidwinter/starr/radarr"
	"time"
)

var logger = utils.GetLogger()

func SearchRadarrMovies(items []model.Item) {
	task := db.LogTaskStart("SearchMissingItemJob")
	r := getRadarrClient()

	var ids []int64
	for _, item := range items {

		logger.Info(fmt.Sprintf("Starting radarr search for %s", item.Title))
		ids = append(ids, item.RadarrId)
		db.MarkItemAsSearched(&item, task)
	}
	resp, err := r.SendCommand(&radarr.CommandRequest{MovieIDs: ids, Name: "MoviesSearch"})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Debug(resp)
	db.LogTaskFinish(task)
}

func ImportRadarrMovies() []model.Item {
	task := db.LogTaskStart("ImportRadarrJob")
	movies := ListRadarrMovies()
	var importedMovies []model.Item

	for _, movie := range movies {
		if movie.HasFile || movie.Status != "released" {
			db.RemoveItem(movie.TmdbID)
			continue
		}
		item := model.Item{
			Name:         movie.Title,
			Available:    movie.HasFile,
			Released:     movie.Status == "released",
			Title:        movie.Title,
			LastSearched: time.Time{},
			TMDBId:       movie.TmdbID,
			RadarrId:     movie.ID,
		}
		logger.Info("Adding", movie)
		importedMovies = append(importedMovies, item)
	}
	db.AddItems(importedMovies)
	db.LogTaskFinish(task)
	return importedMovies
}
func ListRadarrMovies() []*radarr.Movie {

	r := getRadarrClient()

	ctx, cancel := context.WithTimeout(context.TODO(), 180*time.Second)
	defer cancel() // releases resources if slowOperation completes before timeout elapses

	status, err := r.GetSystemStatus()
	if err != nil {
		panic(err)
	}
	logger.Info("Requesting full movie library from radarr, this could take a while")
	m, err := r.GetMovieContext(ctx, 0)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Done requesting full movie library from radarr")

	fmt.Println(status)
	fmt.Println(m)
	return m
}

func getRadarrClient() *radarr.Radarr {
	c := starr.New(viper.GetString("radarr.api_key"), viper.GetString("radarr.url"), 0)
	r := radarr.New(c)
	return r
}
