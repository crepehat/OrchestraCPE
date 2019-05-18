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
	state := heartbeat.State{
		Available:         true,
		MaxOutput:         10,
		MaxOutputDuration: 10,
	}

	heartbeatTicker := time.NewTicker(5 * time.Second)
	configTicker := time.NewTicker(6 * time.Second)

	// heartbeat
	go func() {
		for {
			<-heartbeatTicker.C
			api.SendHeartBeat(state)

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
			api.SendConfig(reqConfig)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", replyState)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func replyState(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Swag."))
}
