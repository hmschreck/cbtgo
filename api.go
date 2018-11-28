package cbtgo

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
)

var username = ""
var authkey = ""

var downloadPath = ""

const API_ROOT = "https://crossbrowsertesting.com/api/v3"

type CBTAPI struct {
	Client resty.Client
}

var APIEndpoints = map[string]string{
	"GetTest":  API_ROOT + "/%s/%d",
	"SetScore": API_ROOT + "/%s/%d",
	"StopTest": API_ROOT + "/%s/%d",
	"TakeSnapshot" : API_ROOT + "/%s/%d/snapshot",
	"SetSnapshot" : API_ROOT + "/%s/%d/snapshot/%s",
	"RecordVideo" : API_ROOT + "/%s/%d/videos",
	"StopVideo" : API_ROOT + "/%s/%d/videos/%s",
}

// Set library-wide download location
func SetDownload(path string) {
	downloadPath = path
}

// Set library-wide authentication options (this is more or less required)
func SetUpAuth(setUsername, setAuthkey string) {
	username = setUsername
	authkey = setAuthkey
}

// Create new API client
func CreateNewAPIClient() *CBTAPI {
	api := new(CBTAPI)
	client := new(resty.Client)
	client.SetBasicAuth(username, authkey)
	client.SetQueryParam("format", "json")
	api.Client = *client
	return api
}

func StopTest(testType string, testID uint64) error {
	api := CreateNewAPIClient()
	endpoint := fmt.Sprintf(APIEndpoints["StopTest"], testType, testID)
	_, err := api.Client.R().Delete(endpoint)
	return err
}

func TakeSnapshot(testType string, testID uint64) (snapshot *Snapshot, errout error) {
	api := CreateNewAPIClient()
	snapshot = new(Snapshot)
	endpoint := fmt.Sprintf(APIEndpoints["TakeSnapshot"], testType, testID)
	response, err := api.Client.R().Post(endpoint)
	if err != nil {
		errout = err
		return
	}
	err = json.Unmarshal(response.Body(), &snapshot)
	if err != nil {
		errout = err
		return
	}
	snapshot.TestID = testID
	return snapshot, nil
}

func SetSnapshot(testType string, testID uint64, hash string) error {
	api := CreateNewAPIClient()
	endpoint := fmt.Sprintf(APIEndpoints["SetSnapshot"], testType, testID, hash)
	_, err := api.Client.R().Put(endpoint)
	return err
}

func RecordVideo(testType string, testID uint64) (video *Video, errout error) {
	api := CreateNewAPIClient()
	video = new(Video)
	endpoint := fmt.Sprintf(APIEndpoints["RecordVideo"], testType, testID)
	response, err := api.Client.R().Post(endpoint)
	if err != nil {
		errout = err
		return
	}
	err = json.Unmarshal(response.Body(), &video)
	if err != nil {
		errout = err
		return
	}
	video.TestID = testID
	video.TestType = testType
	return video, nil
}

func StopVideo(testType string, testID uint64, hash string) (errout error) {
	api := CreateNewAPIClient()
	endpoint := fmt.Sprintf(APIEndpoints["StopVideo"], testType, testID, hash)
	_, err := api.Client.R().Delete(endpoint)
	return err
}