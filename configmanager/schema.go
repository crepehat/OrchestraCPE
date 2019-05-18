package configmanager

type Config struct {
	ObjectId  string  `json:"uid"`
	VillageId string  `json:"vid"`
	Type      string  `json:"type"`
	Values    []Value `json:"values"`
}

type Value struct {
	Format  string      `json:"format"`
	Details interface{} `json:"details"`
}

type CsvDetails struct {
}
