package cbtgo

import (
	"time"
	"fmt"
)

type TestHistory struct {
	Meta struct {
		RecordCount int   `json:"record_count"`
		Source      string `json:"mongodb"`
		Paging      struct {
			Count int `json:"num"`
			Start int `json:"start"`
			NextStart int
		}
		SearchParams struct {
			Active     bool   `json:"active"`
			Resolution string `json:"resolution.name"`
			Start      struct {
				Date time.Time `json:"$gte"`
			} `json:"start_date"`
			End struct {
				Date time.Time `json:"$lte"`
			} `json:"finish_date"`
			BrowserName string `json:"browser.name"`
			BrowserType string `json:"browser.type"`
			OSType      string `json:"os.type"`
			Build       string `json:"caps.build"`
			URL         struct {
				Regex string `json:"$regex"`
			} `json:"url"`
			TestScore string `json:"test_score"`
		} `json:"searchParams"`
	} `json:"meta"`
	SeleniumTests []Test `json:"selenium"`
	LiveTests []Test `json:"livetests"`
	Tests []Test
	TestType string
}

// Merge the  tests of B into A
func (a *TestHistory) Merge(b *TestHistory) {
	a.Tests = append(a.Tests, b.Tests...)
}

func (history *TestHistory) GetAllVideos() {
	count := 0
	for _, test := range history.Tests {
		count += 1
		fmt.Println(count)
		test.GetVideos()
	}
}

func (history *TestHistory) GetAllSnapshots() {
	count := 0
	for _, test := range history.Tests {
		count += 1
		fmt.Println(count)
		test.GetSnapshots()
	}
}

func (history *TestHistory) GetAllHARs() {
	count := 0
	for _, test := range history.Tests {
		count += 1
		fmt.Println(count)
		test.GetHARs()
	}
}

func (history *TestHistory) GetAllPCAPs() {
	count := 0
	for _, test := range history.Tests {
		count += 1
		fmt.Println(count)
		test.GetPCAPs()
	}
}

func (history *TestHistory) WriteAllCommands() {
	count := 0
	for _, test := range history.Tests {
		count += 1
		fmt.Println(count)
		test.WriteCommands()
	}
}

