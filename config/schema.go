package config

type Config struct {
	ObjectId    string  `json:"uid"`
	VillageId   string  `json:"vid"`
	Type        string  `json:"type"`
	Values      []Value `json:"values"`
	Commandable bool    `json:"commandable"`
}

type Value struct {
	Format  string      `json:"format"`
	Details interface{} `json:"details"`
}

type CsvDetails struct {
	FilePath string `json:"file_path"`
}
