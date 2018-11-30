package cbtgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"os"
)

const API_GET_SELENIUM_TEST = "/selenium/%d"
const API_SET_SELENIUM_SCORE = "/selenium/%d"

type Test struct {
	TestID uint64
	LiveTestID 	uint64 `json:"live_test_id,omitempty"`
	SeleniumTestID                   uint64          `json:"selenium_test_id,omitempty"`
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
	NetworkCaptureOn           int            `json:"is_packet_capturing,omitempty"`
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
	NetworkCaptures			[]Network `json:"networks"`
	TestType string
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
	fmt.Println(test.TestID)
	for _, video := range test.Videos {
		fmt.Println("Getting video: ", video.Hash)
		err := video.Get(fmt.Sprintf("%s%s/%d/videos/", DownloadPath, test.TestType, test.TestID))
		if err != nil {
			return err
		}
	}
	return nil
}

func (test Test) GetSnapshots() error {
	fmt.Println(test.TestID)
	for _, snapshot := range test.Snapshots {
		fmt.Println("Getting snapshot: ", snapshot.Hash)
		err := snapshot.Get(fmt.Sprintf("%s%s/%d/snapshots/", DownloadPath, test.TestType, test.TestID))
		if err != nil {
			return err
		}
	}
	return nil
}

func (test Test) GetHARs() error {
	fmt.Println(test.TestID)
	for _, network := range test.NetworkCaptures {
		fmt.Println("Getting HAR: ", network.Hash)
		err := network.Get(fmt.Sprintf("%s%s/%d/hars/", DownloadPath, test.TestType, test.TestID))
		if err != nil {
			return err
		}
	}
	return nil
}

func (test Test) GetPCAPs() error {
	fmt.Println(test.TestID)
	for _, network := range test.NetworkCaptures {
		fmt.Println("Getting HAR: ", network.Hash)
		err := network.GetPcap(fmt.Sprintf("%s%s/%d/pcaps/", DownloadPath, test.TestType, test.TestID))
		if err != nil {
			return err
		}
	}
	return nil
}


func (test *Test) GetVideoData() error {
	api := CreateNewAPIClient()
	endpoint := fmt.Sprintf(APIEndpoints["GetVideoInfo"], test.TestType, test.TestID)
	body, err := api.Client.R().Get(endpoint)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body.Body(), test.Videos)
	if err != nil {
		return err
	}
	return nil
}

func (test *Test) WriteCommands() error {
	if test.TestType != "selenium" {
		return errors.New("Is not a selenium test")
	}
	test, err := test.UpdateCommands()
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmt.Sprintf("%s%s/%d/", DownloadPath, test.TestType, test.TestID), 0755)
	if err != nil {
		return err
	}
	outfile := fmt.Sprintf("%s%s/%d/commands.json", DownloadPath, test.TestType, test.TestID)
	if _, err = os.Stat(outfile); !os.IsNotExist(err) {
		return nil
	}
	file, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Println(test.Commands)
	jsonString, err := json.Marshal(test.Commands)
	fmt.Println(jsonString)
	if err != nil {
		return err
	}
	_, err = file.Write(jsonString)
	if err != nil {
		return err
	}
	return nil
}

func (test *Test) UpdateCommands() (testUpdate *Test, errout error){
	testUpdate, err := GetTest(test.TestType, test.TestID)
	if err != nil {
		errout = err
		return
	}
	return testUpdate, nil
}


