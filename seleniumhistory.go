package cbtgo

import "time"

type SeleniumHistory struct {
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
	SeleniumTests []SeleniumTest `json:"selenium"`
}

// Merge the Selenium tests of SeleniumHistory B into SeleniumHistoryA
func (a *SeleniumHistory) Merge(b *SeleniumHistory) {
	mergeset := make([]SeleniumTest, 0)
	for _, btest := range b.SeleniumTests {
		for i, atest := range a.SeleniumTests {
			if atest.SessionID == btest.SessionID {
				if atest.Active != btest.Active {
					// If btest is the only active one, then A is newer.  Keep it.
					if btest.Active {
						continue
					} else {
						a.SeleniumTests[i] = btest
					}
				}
			} else {
				mergeset = append(mergeset, btest)
			}
		}
	}
	a.SeleniumTests = append(a.SeleniumTests, mergeset...)
}
