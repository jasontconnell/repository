package repository

import (
	"github.com/jasontconnell/conf"
)

type Configuration struct {
	DatabaseServer string `json:"databaseServer"`
	Database string `json:"database"`
	DatabaseUser string `json:"databaseUser"`
	DatabasePassword string `json:"databasePassword"`
}

func LoadConfig(file string) Configuration {
	config := Configuration{}
	conf.LoadConfig(file, &config)
	return config
}