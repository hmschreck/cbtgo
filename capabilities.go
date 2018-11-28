package cbtgo

type Capabilities struct {
	Resolution        string `json:"screenResolution"`
	Name              string `json:"name"`
	RecordSnapshot    bool   `json:"record_snapshot"`
	CBTExample        bool   `json:"cbt_example"`
	Username          string `json:"username"`
	Build             string `json:"build"`
	BrowserName       string `json:"browserName"`
	RecordVideo       bool   `json:"record_video"`
	Platform          string `json:"platform"`
	Version           string `json:"version"`
	CBTScriptID       string `json:"cbt_script_id"`
	CBTScriptVersion  string `json:"cbt_script_version"`
	CBTScriptReplayID string `json:"cbt_script_replay_id"`
}
