package device

import (
	"fmt"
	"net/http"
)

type Command struct {
	Magnitude      int `json:"magnitude"`
	Duration       int `json:"duration"`
	CommandTimeout int `json:"command_timeout"`
}

func (command *Command) ShowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current command:\n%+v\n", *command)
}
