package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/crepehat/OrchestraCPE/api"
	"github.com/crepehat/OrchestraCPE/device"
)

func main() {
	var state device.State
	var currentCommand device.Command
	var err error
	// value
	currentConfig := device.Config{
		ObjectId: "a",
		Type:     "producer",
		// Values:   values,
	}

	state = device.State{
		Available:         true,
		MaxOutput:         10,
		MaxOutputDuration: 10,
	}

	heartbeatTicker := time.NewTicker(1 * time.Second)
	configTicker := time.NewTicker(5 * time.Second)
	stateCheckTicker := time.NewTicker(1 * time.Second)
	stateSetTicket := time.NewTicker(1 * time.Second)

	// check state
	go func() {
		for {
			<-stateCheckTicker.C
			// fmt.Println("Checking local state")
			for _, extractor := range currentConfig.Extractors {
				fmt.Println(extractor)
			}
			for _, insertor := range currentConfig.Inserters {
				fmt.Println(insertor)
			}

		}
	}()

	// set state
	go func() {
		for {
			<-stateSetTicket.C
			// fmt.Println("Setting state")
		}
	}()

	// heartbeat
	go func() {
		for {
			<-heartbeatTicker.C
			currentCommand, err = api.SendHeartBeat(state)
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Println("Got currentCommand back: ", currentCommand)

		}
	}()

	// config checker
	go func() {
		for {
			<-configTicker.C
			reqConfig := device.Config{
				ObjectId:  "69",
				VillageId: "70",
				Type:      "Storage",
				// Values:      values,
				Commandable: true,
			}
			currentConfig, err = api.SyncConfig(reqConfig)
			if err != nil {
				fmt.Println("Error receiving config: ", err.Error())
			}
			fmt.Printf("Received config: %+v\n", currentConfig)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/config", currentConfig.ShowHandler)
	http.HandleFunc("/command", currentCommand.ShowHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
