package main

import (
	"fmt"
	"net/http"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "swag")
	})
	http.ListenAndServe("127.0.0.1:6967", nil)
}
