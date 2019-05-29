package models

type AppConfig struct {
	Host      string                  `json:"host"`
	Databases map[string]DatabaseConf `json:"databases"`
	Usecases  map[string]UsecaseConf  `json:usecases`
}

type DatabaseConf struct {
	URL     string `json:"url"`
	DBName  string `json:"database"`
	TimeOut int    `json:"timeout"`
}

type UsecaseConf struct {
	TimeOut int `json:"timeout"`
}
