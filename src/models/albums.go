package models

const AlbumFileName = "albums.json"

type Album struct {
	Id          string   `json:"id"`
	PhotoCount  string   `json:"photo_count"`
	Url         string   `json:"url"`
	Title       string   `json:"title"`
	Created     string   `json:"created"`
	LastUpdated string   `json:"last_updated"`
	CoverPhoto  string   `json:"cover_photo"`
	Photos      []string `json:"photos"`
}

type Albums []Album

type AlbumList struct {
	Albums Albums `json:"albums"`
}
