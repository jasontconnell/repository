 package repository

import (
	"os"
	"fmt"
	"encoding/json"
)

type Configuration struct {
	Url, Database string
}
var config Configuration = Configuration{}

func LoadConfig(configFile string) Configuration {
	file, err := os.Open(configFile)
 
	if err != nil {
		fmt.Println("Error opening json file ", err, configFile)
	}
 
	decoder := json.NewDecoder(file)

	config := Configuration{}
	err2 := decoder.Decode(&config)
 
	if err2 != nil {
		fmt.Println("error decoding json file", err);
	}

	return config
}