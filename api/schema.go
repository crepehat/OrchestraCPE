package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type API struct {
	Command Command `json:"command"`
	State   State   `json:"state"`
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

func Update(state State) error {
	newApi := API{
		State: state,
	}
	apiString, err := json.Marshal(newApi)
	if err != nil {
		return err
	}
	client.Post("127.0.0.1:6969", "application/json", bytes.NewReader(apiString))
	return nil
}
