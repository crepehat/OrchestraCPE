package device

import (
	"fmt"
	"net/http"
)

type Config struct {
	ObjectId    string      `json:"id"`
	VillageId   string      `json:"vid"`
	Type        string      `json:"type"`
	Extractors  []Extractor `json:"extractor"`
	Inserters   []Inserter  `json:"inserters"`
	Commandable bool        `json:"commandable"`
}

func (config *Config) ShowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current config:\n%+v\n", *config)
	// interact.
}
