package cbtgo

type OSModel struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	Version          string `json:"version"`
	APIName          string `json:"api_name"`
	Device           string `json:"device"`
	RequestedAPIName string `json:"requested_api_name"`
	IconClass        string `json:"icon_class"`
}
