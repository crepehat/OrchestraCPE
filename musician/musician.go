package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/crepehat/OrchestraCPE/heartbeat"

	"github.com/crepehat/OrchestraCPE/api"
	"github.com/crepehat/OrchestraCPE/config"
)

func main() {
	var state heartbeat.State
	var command heartbeat.Command
	var err error
	var currentConfig config.Config
	state = heartbeat.State{
		Available:         true,
		MaxOutput:         10,
		MaxOutputDuration: 10,
	}

	heartbeatTicker := time.NewTicker(5 * time.Second)
	configTicker := time.NewTicker(6 * time.Second)
	stateCheckTicker := time.NewTicker(1 * time.Second)
	stateSetTicket := time.NewTicker(1 * time.Second)

	// check state
	go func() {
		for {
			<-stateCheckTicker.C
			fmt.Println("Checking local state")
			for _, value := range currentConfig.Values {
				fmt.Println(value)
			}
		}
	}()

	// set state
	go func() {
		for {
			<-stateSetTicket.C
			fmt.Println("Setting state")
		}
	}()

	// heartbeat
	go func() {
		for {
			<-heartbeatTicker.C
			command, err = api.SendHeartBeat(state)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(command)

		}
	}()

	// config checker
	go func() {
		for {
			<-configTicker.C
			var values []config.Value
			value := config.Value{
				Format:  "csv",
				Details: "tbd",
			}
			values = append(values, value)
			reqConfig := config.Config{
				ObjectId:    "69",
				VillageId:   "70",
				Type:        "Storage",
				Values:      values,
				Commandable: true,
			}
			api.SyncConfig(reqConfig)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/config", currentConfig.ShowHandler)
	// http.HandleFunc("/config", replyState)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
