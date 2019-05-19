package device

import (
	"fmt"
	"net/http"
)

type Command struct {
	Magnitude      int `json:"magnitude,omitempty"`
	Duration       int `json:"duration,omitempty"`
	CommandTimeout int `json:"command_timeout,omitempty"`
}

func (command *Command) ShowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current command:\n%+v\n", *command)
}
