package cbtgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const API_GET_SELENIUM_TEST = "/selenium/%d"
const API_SET_SELENIUM_SCORE = "/selenium/%d"

type Test struct {
	TestID                   uint64          `json:"selenium_test_id,omitempty"`
	SessionID                string          `json:"selenium_session_id,omitempty"`
	StartDate                time.Time       `json:"start_date,omitempty"`
	FinishDate               time.Time       `json:"finish_date,omitempty"`
	TestScore                string          `json:"test_score,omitempty"`
	ReasonStopped            string          `json:"reason_stopped,omitempty"`
	Active                   bool            `json:"active,omitempty"`
	State                    string          `json:"state,omitempty"`
	VNCPassword              string          `json:"vnc_password,omitempty"`
	VNCPort                  int             `json:"vnc_port,omitempty"`
	Rating                   int             `json:"rating,omitempty"`
	Feedback                 string          `json:"feedback,omitempty"`
	StartupFinishDate        time.Time       `json:"startup_finish_date,omitempty"`
	URL                      string          `json:"url,omitempty"`
	ClientPlatform           string          `json:"client_platform,omitempty"`
	ClientBrowser            string          `json:"client_browser,omitempty"`
	UseCopyrect              bool            `json:"use_copyrect,omitempty"`
	Scale                    string          `json:"scale,omitempty"`
	NetworkCapture           bool            `json:"is_packet_capturing,omitempty"`
	Description              string          `json:"description,omitempty"`
	Tags                     []string        `json:"tags,omitempty"`
	ShowResultUrl            string          `json:"show_result_web_url,omitempty"`
	ShowResultPublicUrl      string          `json:"show_result_public_url,omitempty"`
	DownloadResultsZip       string          `json:"download_results_zip_url,omitempty"`
	DownloadResultsZipPublic string          `json:"download_results_zip_public_url,omitempty"`
	LaunchLiveTestUrl        string          `json:"launch_live_test_url,omitempty"`
	Resolution               ResolutionModel `json:"resolution,omitempty"`
	OperatingSystem          OSModel         `json:"os,omitempty"`
	Browser                  BrowserModel    `json:"browser,omitempty"`
	Videos                   []Video         `json:"videos,omitempty"`
	RequestMethod            string          `json:"requestMethod,omitempty"`
	APIVersion               string          `json:"api_version,omitempty"`
	Snapshots                []Snapshot      `json:"snapshots,omitempty"`
	Archived                 bool            `json:"archived,omitempty"`
	TunnelID                 int             `json:"tunnel_id,omitempty"`
	Commands                 []Command       `json:"commands,omitempty"`
}

func (test *Test) CommandsEmpty() bool {
	if len(test.Commands) == 0 {
		return true
	} else {
		return false
	}
}

func (api *CBTAPI) GetTest(testID int) (selenium_test *Test, errout error) {
	apiCall, err := api.Client.R().Get(APIEndpoints["GetSeleniumTest"])
	if err != nil {
		errout = err
		return
	}
	if apiCall.StatusCode() != 200 {
		errout = errors.New("Received bad response from CBT")
	}
	selenium_test = new(Test)
	body := apiCall.Body()
	err = json.Unmarshal(body, &selenium_test)
	if err != nil {
		errout = err
		return
	}
	if selenium_test.SessionID == "" {
		errout = err
		return
	}
	return selenium_test, nil
}

func (test *Test) SetTest(action string, set string) error {
	api := CreateNewAPIClient()
	endpoint := fmt.Sprintf(APIEndpoints["SetScore"], test.TestID)
	if !(action == "set_score" || action == "set_description") {
		return errors.New("Action must be one of set_score or set_description")
	}
	param := ""
	if action == "set_score" {
		param = "score"
	} else {
		param = "description"
	}
	apiCall, err := api.Client.R().
		SetQueryParam("action", action).
		SetQueryParam(param, "set").
		Put(endpoint)
	if err != nil {
		return err
	}

	if apiCall.StatusCode() != 200 {
		return errors.New("Could not complete set operation")
	}
	return nil
}

func (test *Test) SetScore(score string) error {
	test.TestScore = score
	return test.SetTest("set_score", score)
}
func StopSeleniumTest(testID uint64) error {
	return StopTest("selenium", testID)
}

func (test *Test) Stop() error {
	return StopSeleniumTest(test.TestID)
}
func (test *Test) SetDescription(description string) error {
	test.Description = description
	return test.SetTest("set_description", description)
}

// Get all videos associated with a given test
func (test Test) GetVideos() error {
	for _, video := range test.Videos {
		err := video.Get(fmt.Sprintf("%sselenium/%d/", downloadPath, test.TestID))
		if err != nil {
			return err
		}
	}
	return nil
}


