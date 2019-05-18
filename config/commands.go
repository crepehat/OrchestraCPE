package config

import (
	"fmt"
	"net/http"
)

func (config *Config) ShowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current config:\n%+v\n", *config)
}
