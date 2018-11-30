package cbtgo

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"time"
)

var Username = ""
var Authkey = ""

var DownloadPath = ""
const MAX_RESULTS = 2500

const API_ROOT = "https://crossbrowsertesting.com/api/v3"

type CBTAPI struct {
	Client resty.Client
}

var APIEndpoints = map[string]string{
	"GetTestHistory" : API_ROOT + "/%s",
	"GetTest":  API_ROOT + "/%s/%d",
	"SetScore": API_ROOT + "/%s/%d",
	"StopTest": API_ROOT + "/%s/%d",
	"TakeSnapshot" : API_ROOT + "/%s/%d/snapshot",
	"SetSnapshot" : API_ROOT + "/%s/%d/snapshot/%s",
	"RecordVideo" : API_ROOT + "/%s/%d/videos",
	"StopVideo" : API_ROOT + "/%s/%d/videos/%s",
	"SetVideoDescription" : API_ROOT + "/%s/%d/videos/%s",
	"RecordNetworkPackets": API_ROOT + "/%s/%d/networks",
	"StopNetworkRecord": API_ROOT + "/%s/%d/networks/%s",
	"SetNetworkDescription" : API_ROOT + "/%s/%d/networks/%s",
	"GetVideoInfo": API_ROOT + "/%s/%d/videos",
}

// Set library-wide download location
func SetDownload(path string) {
	DownloadPath = path
}

// Set library-wide authentication options (this is more or less required)
func SetUpAuth(setUsername, setAuthkey string) {
	Username = setUsername
	Authkey = setAuthkey
}

// Create new API client
func CreateNewAPIClient() *CBTAPI {
	api := new(CBTAPI)
	client := resty.New()
	client = client.SetBasicAuth(Username, Authkey)
	client = client.SetQueryParam("format", "json")
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

func SetVideoDescription(testType string, testID uint64, hash string, description string) (error) {
	api := CreateNewAPIClient()
	endpoint := fmt.Sprintf(APIEndpoints["SetVideoDescription"], testType, testID, hash)
	_, err := api.Client.SetQueryParam("description", description).R().Put(endpoint)
	return err
}

func RecordNetworkPackets(testType string, testID uint64) (network *Network, errout error) {
	network = new(Network)
	api := CreateNewAPIClient()
	endpoint := fmt.Sprint(APIEndpoints["RecordNetworkPackets"], testType, testID)
	response, err := api.Client.R().Post(endpoint)
	if err != nil {
		errout = err
		return
	}
	errout = json.Unmarshal(response.Body(), &network)
	network.TestID = testID
	network.TestType = testType
	return
}

func StopNetworkPackets(testType string, testID uint64, hash string) error {
	api := CreateNewAPIClient()
	endpoint := fmt.Sprintf(APIEndpoints["StopNetworkRecord"], testType, testID, hash)
	_, err := api.Client.R().Delete(endpoint)
	return err
}

func SetNetworkDescription(testType string, testID uint64, hash string, description string) error {
	api := CreateNewAPIClient()
	endpoint := fmt.Sprintf(APIEndpoints["SetNetworkDescription"], testType, testID, hash)
	_, err := api.Client.SetQueryParam("description", description).R().Put(endpoint)
	return err
}

func GetTestHistory(testType string, num int, start int) (testhistory *TestHistory, errout error) {
	api := CreateNewAPIClient()
	testhistory = new(TestHistory)
	stringNum := fmt.Sprintf("%d", num)
	stringStart := fmt.Sprintf("%d", start)
	endpoint := fmt.Sprintf(APIEndpoints["GetTestHistory"], testType)
	response, err := api.Client.R().
		SetQueryParam("num", stringNum).
		SetQueryParam("start", stringStart).
		Get(endpoint)
	if err != nil {
		errout = err
		return
	}
	err = json.Unmarshal(response.Body(), &testhistory)

	if testhistory.SeleniumTests == nil {
		testhistory.Tests = testhistory.LiveTests
		testhistory.LiveTests = nil
	} else if testhistory.LiveTests == nil {
		testhistory.Tests = testhistory.SeleniumTests
		testhistory.SeleniumTests = nil
	}	
	for i, test := range testhistory.Tests {
		test.TestType = testType
		if test.SeleniumTestID ==  0{
			test.TestID = test.LiveTestID
		} else if test.LiveTestID == 0 {
			test.TestID = test.SeleniumTestID
		}
		testhistory.Tests[i] = test
	}
	return testhistory, err
}

func GetCompleteTestHistory(testType string) (testhistoryfull *TestHistory, errout error) {
	start := 0
	num := 25
	currentCount := 1
	testhistory, err := GetTestHistory(testType, num, start)
	if err != nil {
		errout = err
		return
	}
	currentCount = testhistory.Meta.RecordCount
	testhistoryfull = testhistory
	start = start + num
	for {

		fmt.Println(len(testhistoryfull.Tests))
		if start >= currentCount || start >= MAX_RESULTS {
			break
		}
		newhistory, err := GetTestHistory(testType, num, start)
		if err != nil {
			errout = err
			return
		}
		testhistoryfull.Merge(newhistory)
		start = start + num
		time.Sleep(1 * time.Millisecond)
	}
	return
}

func GetTest(testType string, testID uint64) (test *Test, errout error) {
	api := CreateNewAPIClient()
	test = new(Test)
	endpoint := fmt.Sprintf(APIEndpoints["GetTest"], testType, testID)
	response, err := api.Client.R().Get(endpoint)
	if err != nil {
		errout = err
		return
	}
	err = json.Unmarshal(response.Body(), &test)
	if err != nil {
		errout = err
		return
	}
	test.TestID = testID
	test.TestType = testType
	return test, nil
}