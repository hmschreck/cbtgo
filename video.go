package cbtgo

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Video struct {
	Hash                string    `json:"hash,omitempty"`
	DateAdded           time.Time `json:"date_added,omitempty"`
	Description         string    `json:"description,omitempty"`
	Tags                []string  `json:"tags,omitempty"`
	Finished            bool      `json:"is_finished,omitempty"`
	ShowResultUrl       string    `json:"show_result_web_url,omitempty"`
	ShowResultPublicUrl string    `json:"show_result_public_url,omitempty"`
	Video               string    `json:"video,omitempty"`
	Image               string    `json:"image,omitempty"`
	Thumbnail           string    `json:"thumbnail_image,omitempty"`
	TestID uint64
	TestType string
}

// Download a video
func (video *Video) Get(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	filepath := fmt.Sprintf("%s%s.mp4", path, video.Hash)
	if _, err = os.Stat(filepath); !os.IsNotExist(err) {
		// File already exists; return
		return nil
	}
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	file, err := http.Get(video.Video)
	if err != nil {
		return err
	}
	defer file.Body.Close()
	_, err = io.Copy(out, file.Body)
	if err != nil {
		return err
	}
	return nil
}

func (video *Video) Stop() error {
	return StopVideo(video.TestType, video.TestID, video.Hash)
}
