package device

import (
	"fmt"
	"net/http"
)

type Config struct {
	ObjectId    string      `json:"id"`
	VillageId   string      `json:"vid"`
	Type        string      `json:"type,omitempty"`
	Legacy      bool        `json:"legacy,omitempty"`
	Extractors  []Extractor `json:"extractor,omitempty"`
	Inserters   []Inserter  `json:"inserters,omitempty"`
	Commandable bool        `json:"commandable,omitempty"`
}

func (config *Config) ShowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current config:\n%+v\n", *config)
	// interact.
}
