package cbtgo

import (
	"time"
	"fmt"
	"os"
	"net/http"
	"encoding/json"
	"io"
)
const MAX_SCREENSHOTS = 500



func GetScreenshotHistory(num int, start int) (testhistory *ScreenshotHistory, errout error) {
	api := CreateNewAPIClient()
	testhistory = new(ScreenshotHistory)
	endpoint := fmt.Sprintf(APIEndpoints["GetTestHistory"], "screenshots")
	numString := fmt.Sprintf("%d", num)
	startString := fmt.Sprintf("%d", start)
	body, err := api.Client.SetQueryParam("num", numString).SetQueryParam("start", startString).R().Get(endpoint)
	if err != nil {
		errout = err
		return
	}
	errout = json.Unmarshal(body.Body(), &testhistory)
	return
}

func (a *ScreenshotHistory) Merge(b *ScreenshotHistory) error {
	a.ScreenshotTests = append(a.ScreenshotTests, b.ScreenshotTests...)
	return nil
}

func GetFullScreenshotHistory() (testhistory *ScreenshotHistory, errout error) {
	start := 0
	num := 10
	currentCount := 1
	testhistory, err := GetScreenshotHistory(num, start)
	if err != nil {
		errout = err
		return
	}
	currentCount = testhistory.Meta.RecordCount
	start = start + num
	for {
		if start > currentCount || start >= MAX_SCREENSHOTS {
			break
		}
		newhistory, err := GetScreenshotHistory(num, start)
		if err != nil {
			errout = err
			return
		}
		testhistory.Merge(newhistory)
		start = start+num
		time.Sleep(3000 * time.Millisecond)
	}
	return
}

func (screenshothistory *ScreenshotHistory) GetAll() error {
	for _, screenshottest := range screenshothistory.ScreenshotTests {
		err := screenshottest.Get()
		if err != nil {
			return err
		}
	}
	return nil
}

func (test *ScreenshotTest) Get() error{
	for _, version := range test.Versions {
		filepath := fmt.Sprintf("%sscreenshots/%d/%d", DownloadPath, test.ScreenshotTestID, version.VersionID)
		for _, result := range version.Results {
			err := os.MkdirAll(filepath, 0755)
			chromelessOut := fmt.Sprintf("%s/%s-chromeless.png", filepath, result.Images.Hash)
			windowedOut := fmt.Sprintf("%s/%s-windowed.png", filepath, result.Images.Hash)
			fullscreenOut := fmt.Sprintf("%s/%s-fullscreen.png", filepath, result.Images.Hash)
			if _, err := os.Stat(fullscreenOut); !os.IsNotExist(err) {
				continue
			}
			chromeFile, err := os.Create(chromelessOut)
			windowedFile, err := os.Create(windowedOut)
			fullscreenFile, err := os.Create(fullscreenOut)
			defer chromeFile.Close()
			defer windowedFile.Close()
			defer fullscreenFile.Close()
			chromeDownload, err := http.Get(result.Images.Chromeless)
			defer chromeDownload.Body.Close()
			windowedDownload, err := http.Get(result.Images.Windowed)
			defer windowedDownload.Body.Close()
			fullscreenDownload, err := http.Get(result.Images.Fullpage)
			defer fullscreenDownload.Body.Close()
			_, err = io.Copy(chromeFile, chromeDownload.Body)
			_, err = io.Copy(windowedFile, windowedDownload.Body)
			_, err = io.Copy(fullscreenFile, fullscreenDownload.Body)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
