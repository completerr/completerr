package services

import (
	"completerr/db"
	"completerr/model"
	"context"
	"github.com/craigjmidwinter/starr"
	"github.com/craigjmidwinter/starr/sonarr"
	"github.com/spf13/viper"
	"net/url"
	"strconv"
	"time"
)

type EpisodeSearchCommandRequest struct {
	Name       string  `json:"name"`
	Files      []int64 `json:"files,omitempty"` // RenameFiles only
	SeriesIDs  []int64 `json:"seriesIds,omitempty"`
	SeriesID   int64   `json:"seriesId,omitempty"`
	EpisodeIDs int64   `json:"episodeIds,omitempty"`
}

func SearchSonarrEpisode(items []model.TvItem) {
	task := db.LogTaskStart("SearchMissingTvItemJob")
	r := getSonarrClient()

	ctx, cancel := context.WithTimeout(context.TODO(), 180*time.Second)
	defer cancel() // releases resources if slowOperation completes before timeout elapses
	var ids []int64

	for _, item := range items {
		ids = append(ids, item.SonarrId)
		db.MarkTvItemAsSearched(&item, task)
	}
	resp, err := r.SendCommandContext(ctx, &sonarr.CommandRequest{EpisodeIDs: ids, Name: "EpisodeSearch"})

	if err != nil {
		logger.Error(err)
		return
	}
	logger.Debug(resp)
	db.LogTaskFinish(task)
}

func ImportSonarrEpisodes() []model.TvItem {
	task := db.LogTaskStart("ImportSonarrJob")
	db.MarkAllTVItemsAsDeleted()
	episodes := ListSonarrMovies()
	var importedEpisodes []model.TvItem

	for _, episode := range episodes {
		item := model.TvItem{
			Name:                  episode.Title,
			Title:                 episode.Title,
			Available:             episode.HasFile,
			LastSearched:          time.Time{},
			SonarrId:              episode.ID,
			SonarrSeriesId:        episode.SeriesID,
			SeriesTitle:           episode.Series.Title,
			Season:                episode.SeasonNumber,
			SeasonEpisodeNumber:   episode.EpisodeNumber,
			AbsoluteEpisodeNumber: episode.AbsoluteEpisodeNumber,
		}
		importedEpisodes = append(importedEpisodes, item)
		logger.Info("Adding", episode)
	}
	db.AddTVItems(importedEpisodes)
	db.LogTaskFinish(task)
	return importedEpisodes
}

type Wanted struct {
	Page          int               `json:"page"`
	PageSize      int               `json:"pageSize"`
	SortKey       string            `json:"sortKey"`
	SortDirection string            `json:"sortDirection"`
	TotalRecords  int               `json:"totalRecords"`
	Records       []*sonarr.Episode `json:"records"`
}

func ListSonarrMovies() []*sonarr.Episode {
	r := getSonarrClient()

	ctx, cancel := context.WithTimeout(context.TODO(), 180*time.Second)
	defer cancel() // releases resources if slowOperation completes before timeout elapses

	resp := Wanted{}
	logger.Info("Requesting full movie library from sonarr, this could take a while")
	pageSize := -1
	var records []*sonarr.Episode
	err := r.GetInto(
		ctx,
		starr.Request{URI: "/wanted/missing",
			Query: url.Values{"page": {"1"}, "pageSize": {strconv.Itoa(pageSize)}, "sortDirection": {"descending"}, "monitored": {"true"}, "sortKey": {"airDateUtc"}},
		},
		&resp,
	)
	if err != nil {
		logger.Error(err)
	}
	records = append(records, resp.Records...)
	logger.Debug(resp)
	logger.Info("Done requesting full movie library from sonarr")

	return records
}

func getSonarrClient() *sonarr.Sonarr {
	c := starr.New(viper.GetString("sonarr.api_key"), viper.GetString("sonarr.url"), 0)
	r := sonarr.New(c)
	return r
}
