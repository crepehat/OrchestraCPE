package api

import (
	"net/http"
	"time"
)

type HeartBeat struct {
	DeviceId string  `json:"device_id"`
	Command  Command `json:"command"`
	State    State   `json:"state"`
}

type Command struct {
	Magnitude int `json:"magnitude"`
	Duration  int `json:"duration"`
}

type State struct {
	Commandable       bool `json:"commandable"`
	MaxOutput         int  `json:"max_output"`
	MaxOutputDuration int  `json:"max_output_duration"`
}

var client http.Client

func init() {
	client = http.Client{
		Timeout: 30 * time.Second,
	}
}
