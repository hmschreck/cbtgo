package cbtgo

type ResolutionModel struct {
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	DesktopWidth  int    `json:"desktop_width"`
	DesktopHeight int    `json:"desktop_height"`
	Name          string `json:"name"`
}
