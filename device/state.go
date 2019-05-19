package device

type State struct {
	Available         bool    `json:"available,omitempty"`
	MaxOutput         int     `json:"max_output,omitempty"`
	MaxOutputDuration int     `json:"max_output_duration,omitempty"`
	CurrentCommand    Command `json:"current_command,omitempty"`
	Energised         bool    `json:"energised,omitempty"`
	CurrentOutput     int     `json:"current_output,omitempty"`
}
