package device

type State struct {
	Available         bool    `json:"available"`
	MaxOutput         int     `json:"max_output"`
	MaxOutputDuration int     `json:"max_output_duration"`
	CurrentCommand    Command `json:"current_command"`
	Energised         bool    `json:"energised"`
	CurrentOutput     int     `json:"current_output"`
}
