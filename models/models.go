package models

import "time"

// struct that represents the netflix catalog table
type MediaInfo struct {
	ID          int
	ShowID      string
	Type        string
	Title       string
	Director    string
	Country     string
	DateAdded   time.Time
	ReleaseYear int
	Rating      string
	Duration    string
	Category    string
	is_deleted  bool
	Favorited   interface{}
}

// struct that represents the watched_media table
type WatchListInfo struct {
	Media   MediaInfo
	ID      int
	Watched bool
}
