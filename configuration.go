package repository

import (
	"conf"
)

type Configuration struct {
	DatabaseServer string `json:"databaseServer"`
	Database string `json:"database"`
}

func LoadConfig(file string) Configuration {
	config := Configuration{}
	conf.LoadConfig(file, &config)
	return config
}