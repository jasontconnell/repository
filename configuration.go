 package repository

import (
	"conf"
)

type Configuration struct {
	Url string `json:"Url"`
	Database string `json:"Database"`
}

func (config *Configuration) LoadConfig(file string){
	conf.LoadConfig(file, config)
}