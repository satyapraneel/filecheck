package services

import (
	"encoding/json"
	"filecheck/config"
	"io/ioutil"
)

type App struct {
	DAConfig  config.DAConfig
	Scheduler config.Scheduler
}

func GetConfigurationDetails() config.DAConfig {
	path := config.GetFileConfig().JsonFilePath
	jsonFile, _ := ioutil.ReadFile(path)
	var data config.DAConfig
	err := json.Unmarshal(jsonFile, &data)
	if err != nil {
		panic(err)
	}
	return data
}
