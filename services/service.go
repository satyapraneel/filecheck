package services

import (
	"encoding/json"
	"filecheck/config"
	"io/ioutil"
)

type App struct {
	DAConfig config.DAConfig
}

func GetConfigurationDetails() config.DAConfig {
	path := config.GetFileConfig().JsonFilePath
	jsonFile, _ := ioutil.ReadFile(path)
	var data config.DAConfig
	err := json.Unmarshal(jsonFile, &data)
	if err != nil {
		panic(path + " Json file not found")
	}
	return data
}
