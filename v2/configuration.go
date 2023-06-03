package repository

import (
	"github.com/jasontconnell/conf"
)

type Configuration struct {
	ConnectionString string `json:"connectionString"`
	Database         string `json:"database"`
}

func LoadConfig(file string) Configuration {
	config := Configuration{}
	conf.LoadConfig(file, &config)
	return config
}
