package cbtgo

import "time"

type Snapshot struct {
	Hash                string    `json:"hash"`
	DateAdded           time.Time `json:"date_added"`
	Description         string    `json:"description"`
	Tags                []string  `json:"tags"`
	ShowResultUrl       string    `json:"show_result_web_url"`
	ShowResultPublicUrl string    `json:"show_result_public_url"`
	Image               string    `json:"image"`
	Thumbnail           string    `json:"thumbnail_image"`
	TestID uint64
}

