package api

import (
	"bytes"
	"encoding/json"
)

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
