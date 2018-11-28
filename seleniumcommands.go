package cbtgo

import "time"

type Command struct {
	Body         string    `json:"body,omitempty"`
	Method       string    `json:"method,omitempty"`
	Path         string    `json:"path,omitempty"`
	DateIssued   time.Time `json:"date_issued,omitempty"`
	Hash         string    `json:"hash,omitempty"`
	ResponseCode int       `json:"response_code,omitempty"`
	ResponseBody string    `json:"response_body,omitempty"`
	StepNumber   int       `json:"step_number,omitempty"`
}
