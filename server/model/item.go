package model

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
	gorm.Model
	ID           int64     `json:"ID"`
	Name         string    `json:"name"`
	Available    bool      `json:"available"`
	Released     bool      `json:"released"`
	Title        string    `json:"title"`
	LastSearched time.Time `json:"last_searched"`
	TMDBId       int64     `json:"tmdb_id" gorm:"uniqueIndex;not null"`
	RadarrId     int64     `json:"radarr_id" gorm:"uniqueIndex;not null"`
	SearchCount  int       `json:"search_count" gorm:"default: 0;not null"`
}
type TvItem struct {
	gorm.Model
	ID                    int64     `json:"ID"`
	Name                  string    `json:"name"`
	Available             bool      `json:"available"`
	Title                 string    `json:"title"`
	LastSearched          time.Time `json:"last_searched"`
	SonarrId              int64     `json:"sonarr_id" gorm:"uniqueIndex;not null"`
	SearchCount           int       `json:"search_count" gorm:"default: 0;not null"`
	SonarrSeriesId        int64     `json:"sonarr_series_id"`
	SeriesTitle           string    `json:"series_title"`
	Season                int64     `json:"season"`
	SeasonEpisodeNumber   int64     `json:"season_episode_number"`
	AbsoluteEpisodeNumber int64     `json:"absolute_episode_number"`
}
type Task struct {
	gorm.Model
	ID       int64     `json:"ID"`
	Type     string    `json:"type"`
	Status   string    `json:"status"`
	Started  time.Time `json:"started"`
	Finished time.Time `json:"finished"`
}
type SearchRecord struct {
	gorm.Model
	ID       int64  `json:"ID"`
	TaskID   int64  `json:"task_id"`
	Task     Task   `json:"task"`
	TvItemID int64  `json:"tv_item_id,omitempty"`
	TvItem   TvItem `json:"tv_item,omitempty"`
	Item     Item   `json:"item,omitempty"`
	ItemID   int64  `json:"item_id,omitempty"`
}
