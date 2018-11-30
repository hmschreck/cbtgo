package cbtgo

import "testing"

const TestUser = "test"
const TestKey = "pass"
const TestDownloadPath = "./"
func TestSetUpAuth(t *testing.T) {
	SetUpAuth(TestUser, TestKey)
	if !(Username == TestUser && Authkey == TestKey) {
		t.Errorf("Could not set authentication")
	}
}

func TestSetDownload(t *testing.T) {
	SetDownload(TestDownloadPath)
	if DownloadPath != TestDownloadPath {
		t.Errorf("Could not set downloadpath")
	}
}
