package cbtgo

import "time"

type TestHistory struct {
	Meta struct {
		RecordCount uint   `json:"record_count"`
		Source      string `json:"mongodb"`
		Paging      struct {
			Count int `json:"num"`
			Start int `json:"start"`
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
	Tests []Test `json:"selenium"`
}

// Merge the  tests of B into A
func (a *TestHistory) Merge(b *TestHistory) {
	mergeset := make([]Test, 0)
	for _, btest := range b.Tests {
		for i, atest := range a.Tests {
			if atest.SessionID == btest.SessionID {
				if atest.Active != btest.Active {
					// If btest is the only active one, then A is newer.  Keep it.
					if btest.Active {
						continue
					} else {
						a.Tests[i] = btest
					}
				}
			} else {
				mergeset = append(mergeset, btest)
			}
		}
	}
	a.Tests = append(a.Tests, mergeset...)
}
