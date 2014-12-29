package repository

import (
	"conf"
)

type Configuration struct {
	Url string `json:"url"`
	Database string `json:"database"`
}

func LoadConfig(file string) Configuration {
	config := Configuration{}
	conf.LoadConfig(file, &config)
	return config
}