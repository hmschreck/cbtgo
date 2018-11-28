package cbtgo

import (
	"github.com/pkg/errors"
	"time"
)

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
	TestID uint64
	TestType string
}


func (network *Network) Stop() error{
	if network.TestType == "" || network.TestID == 0 {
		return errors.New("Do not have all the necessary information to stop the record.")
	}
	return StopNetworkPackets(network.TestType, network.TestID, network.Hash)
}

func (network *Network) SetDescription(description string) error {
	network.Description = description
	if network.TestType == "" || network.TestID == 0 {
		return errors.New("Do not have all the necessary information to stop the record.")
	}
	return SetNetworkDescription(network.TestType, network.TestID, network.Hash, description)
}
