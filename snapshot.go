package cbtgo

import (
	"time"
	"os"
	"io"
	"fmt"
	"net/http"
)

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

func (snapshot *Snapshot) Get(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	filepath := fmt.Sprintf("%s%s.png", path, snapshot.Hash)
	if _, err = os.Stat(filepath); !os.IsNotExist(err) {
		// File already exists; return
		return nil
	}
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	file, err := http.Get(snapshot.Image)
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

