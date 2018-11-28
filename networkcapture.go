package cbtgo

import "time"

type Network struct {
	Hash                string    `json:"hash"`
	DateAdded           time.Time `json:"date_added"`
	Description         string    `json:"description"`
	Tags                []string  `json:"tags"`
	Finished            bool      `json:"is_finished"`
	ShowResultUrl       string    `json:"show_result_web_url"`
	ShowResultPublicUrl string    `json:"show_result_public_url"`
	PCAP                string    `json:"pcap"`
	HAR                 string    `json:"har"`
}
