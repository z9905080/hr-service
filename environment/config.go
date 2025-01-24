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
