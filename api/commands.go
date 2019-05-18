package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/crepehat/OrchestraCPE/device"
)

func SendHeartBeat(state device.State) (device.Command, error) {
	heartBeat := HeartBeat{
		DeviceId: "10",
		State:    state,
	}
	apiString, err := json.Marshal(heartBeat)
	if err != nil {
		return heartBeat.Command, err
	}
	fmt.Printf("Sending state: %s\n", apiString)
	resp, err := client.Post(heartbeatApi, "application/json", bytes.NewReader(apiString))
	if err != nil {
		fmt.Println(err)
		return heartBeat.Command, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(bodyBytes))
	var heartBeatReceived HeartBeat
	err = json.Unmarshal(bodyBytes, &heartBeatReceived)
	fmt.Printf("Received command: %+v\n", heartBeatReceived.Command)

	return heartBeatReceived.Command, nil
}

func SyncConfig(reqConfig device.Config) (device.Config, error) {
	var retConfig device.Config
	apiString, err := json.Marshal(reqConfig)
	if err != nil {
		return retConfig, err
	}
	fmt.Printf("Sending config: %s\n", apiString)
	resp, err := client.Post(configApi, "application/json", bytes.NewReader(apiString))
	if err != nil {
		fmt.Println(err)
		return retConfig, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(bodyBytes))
	var returnedConfig device.Config
	err = json.Unmarshal(bodyBytes, &returnedConfig)
	fmt.Printf("Received config: %+v\n", returnedConfig)

	return retConfig, nil
}
