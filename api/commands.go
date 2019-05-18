package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/crepehat/OrchestraCPE/config"
	"github.com/crepehat/OrchestraCPE/heartbeat"
)

func SendHeartBeat(state heartbeat.State) error {
	heartBeat := heartbeat.HeartBeat{
		DeviceId: "10",
		State:    state,
	}
	apiString, err := json.Marshal(heartBeat)
	if err != nil {
		return err
	}
	fmt.Printf("Sending state: %s\n", apiString)
	resp, err := client.Post(heartbeatApi, "application/json", bytes.NewReader(apiString))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	var commandReceived heartbeat.Command
	err = json.Unmarshal(bodyBytes, &commandReceived)
	fmt.Printf("Received command: %+v\n", commandReceived)

	return nil
}

func SendConfig(reqConfig config.Config) error {

	apiString, err := json.Marshal(reqConfig)
	if err != nil {
		return err
	}
	fmt.Printf("Sending config: %s\n", apiString)
	resp, err := client.Post(configApi, "application/json", bytes.NewReader(apiString))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	var returnedConfig config.Config
	err = json.Unmarshal(bodyBytes, &returnedConfig)
	fmt.Printf("Received config: %+v\n", returnedConfig)

	return nil
}
