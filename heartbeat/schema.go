package heartbeat

type HeartBeat struct {
	DeviceId string  `json:"device_id"`
	Command  Command `json:"command"`
	State    State   `json:"state"`
}

type Command struct {
	Magnitude      int `json:"magnitude"`
	Duration       int `json:"duration"`
	CommandTimeout int `json:"command_timeout"`
}

type State struct {
	Available         bool    `json:"available"`
	MaxOutput         int     `json:"max_output"`
	MaxOutputDuration int     `json:"max_output_duration"`
	CurrentCommand    Command `json:"current_command"`
}
