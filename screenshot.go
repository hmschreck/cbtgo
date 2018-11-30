package cbtgo

import (
	"time"
)

type ScreenshotHistory struct {
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
	ScreenshotTests []ScreenshotTest `json:"screenshots"`
}

type Version struct {
	Rating string `json:"rating"`
	Feedback string `json:"feedback"`
	ComparisonRating string `json:"comparison_rating"`
	ComparisonFeedback string `json:"comparison_feedback"`
	TunnelID int `json:"tunnel_id"`
	ApplitoolsBatchID string `json:"applitools_batch_id"`
	VersionID int `json:"version_id"`
	StartDate time.Time `json:"start_date"`
	VersionHash string `json:"version_hash"`
	RequestMethod string `json:"requestMethod"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	ResultCount struct {
		Total int `json:"total"`
		Running int `json:"running"`
		Successful int `json:"successful"`
		Failed int `json:"failed"`
		Cancelled int `json:"cancelled"`
	} `json:"result_count"`
	Active bool `json:"active"`
	ShowResultUrl            string          `json:"show_result_web_url,omitempty"`
	ShowResultPublicUrl      string          `json:"show_result_public_url,omitempty"`
	DownloadResultsZip       string          `json:"download_results_zip_url,omitempty"`
	DownloadResultsZipPublic string          `json:"download_results_zip_public_url,omitempty"`
	ShowComparisonsWebUrl string `json:"show_comparisons_web_url"`
	ShowComparisonsPublicUrl string `json:"show_comparisons_public_url"`
	Results []ScreenshotResult `json:"results"`
	TestSchedule string `json:"test_schedule"`
}

type ScreenshotResult struct {
	FinishDate time.Time `json:"finish_date"`
	Flagged int `json:"flagged"`
	FlaggedReason string `json:"flagged_reason"`
	FlaggedDetails string `json:"flagged_details"`
	ResultHash string `json:"result_hash"`
	CapturedDom bool `json:"captured_dom"`
	ApplitoolsSessionID string `json:"applitools_session_id"`
	ApplitoolsAccessToken string `json:"applitools_access_token"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	State string `json:"state"`
	Successful bool `json:"successful"`
	Resolution ResolutionModel `json:"resolution"`
	ResultID int `json:"result_id"`
	InitializedDate time.Time `json:"initialized_date"`
	StartDate time.Time `json:"start_date"`
	OS OSModel `json:"os"`
	Browser BrowserModel `json:"browser"`
	Images struct {
		Hash string `json:"hash"`
		Windowed string `json:"windowed"`
		Chromeless string `json:"chromeless"`
		Fullpage string `json:"fullpage"`
	} `json:"images"`
	LaunchLiveTestURL string `json:"launch_live_test_url"`
	PageSource string `json:"page_source"`
	DomSource string `json:"dom_source"`
}

type ScreenshotTest struct {
	ScreenshotTestID int `json:"screenshot_test_id"`
	Url string `json:"url"`
	Created time.Time `json:"created_date"`
	Archived bool `json:"archived"`
	VersionCount int `json:"version_count"`
	APIVersion string `json:"api_version"`
	Options ScreenshotOption `json:"options"`
	Versions []Version `json:"versions"`
}

type ScreenshotOption struct {
	Delay int `json:"delay"`
	HideFixedElements bool `json:"hide_fixed_elements"`
	UseBasicAuth bool `json:"use_basic_auth"`
	BasicUsername string `json:"basic_username"`
	UseLoginProfile bool `json:"use_login_profile"`
	UseSeleniumScript bool `json:"use_selenium_script"`
	LoginProfile string `json:"login_profile"`
	SeleniumScript string `json:"selenium_script"`
	SendEmail bool `json:"send_email"`
	EmailLIst string `json:"email_list"`
	SendToApplitools bool `json:"send_to_applitools"`
	ApplitoolsTestName bool `json:"applitools_test_name"`
}