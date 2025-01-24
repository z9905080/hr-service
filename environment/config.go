package environment

import (
	"encoding/json"
	"os"
)

type Config struct {
	Http struct {
		Port string `json:"port"`
	} `json:"http"`
	LogLevel string `json:"log_level"`
	DB       struct {
		User           string `json:"user"`
		Password       string `json:"password"`
		InstanceName   string `json:"instance_name"`
		Database       string `json:"database"`
		ConnectTimeout string `json:"connect_timeout"`
		ReadTimeout    string `json:"read_timeout"`
		WriteTimeout   string `json:"write_timeout"`
	} `json:"db"`
}

type ConfigPathType string

var CnfPath ConfigPathType

func StaticPath() ConfigPathType {
	return CnfPath
}

func NewConfig(path ConfigPathType) Config {
	// read from filepath
	filePath := string(path)
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
