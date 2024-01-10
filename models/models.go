package models

import "time"

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
}

type MediaInfoWatch struct {
	ShowID  string
	Watched bool
}

type WatchListInfo struct {
	Media   MediaInfo
	ID      int
	Watched bool
}
