package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/crepehat/OrchestraCPE/api"
	"github.com/crepehat/OrchestraCPE/inputs"
)

func main() {
	state := api.State{
		Commandable:       true,
		MaxOutput:         10,
		MaxOutputDuration: 10,
	}

	updateTicker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			<-updateTicker.C
			api.Update(state)

		}
	}()

	go func() {
		for {
			<-updateTicker.C
			inputs.CsvGetValue("swag", 2)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "swag")
	})
	http.ListenAndServe("127.0.0.1:6967", nil)
}
