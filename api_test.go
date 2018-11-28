package cbtgo

import "testing"

const TestUser = "test"
const TestKey = "pass"
const TestDownloadPath = "./"
func TestSetUpAuth(t *testing.T) {
	SetUpAuth(TestUser, TestKey)
	if !(username == TestUser && authkey == TestKey) {
		t.Errorf("Could not set authentication")
	}
}

func TestSetDownload(t *testing.T) {
	SetDownload(TestDownloadPath)
	if downloadPath != TestDownloadPath {
		t.Errorf("Could not set downloadpath")
	}
}
