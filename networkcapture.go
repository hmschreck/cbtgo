package cbtgo

import (
	"github.com/pkg/errors"
	"time"
	"os"
	"net/http"
	"io"
	"fmt"
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

func (network *Network) Get(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	filepath := fmt.Sprintf("%s%s.har", path, network.Hash)
	if _, err = os.Stat(filepath); !os.IsNotExist(err) {
		// File already exists; return
		return nil
	}
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	file, err := http.Get(network.HAR)
	if err != nil {
		return err
	}
	defer file.Body.Close()
	_, err = io.Copy(out, file.Body)
	if err != nil {
		return err
	}
	return nil
}
func (network *Network) GetPcap(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	filepath := fmt.Sprintf("%s%s.pcap", path, network.Hash)
	if _, err = os.Stat(filepath); !os.IsNotExist(err) {
		// File already exists; return
		return nil
	}
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	file, err := http.Get(network.PCAP)
	if err != nil {
		return err
	}
	defer file.Body.Close()
	_, err = io.Copy(out, file.Body)
	if err != nil {
		return err
	}
	return nil
}
