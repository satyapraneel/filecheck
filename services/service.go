package services

import (
	"encoding/json"
	"filecheck/config"
	"filecheck/jsons"
	"io/ioutil"
)

type App struct {
	DAStruct jsons.DAStruct
}

func GetConfigurationDetails() jsons.DAStruct {
	path := config.GetFileConfig().JsonFilePath
	jsonFile, _ := ioutil.ReadFile(path)
	var data jsons.DAStruct
	err := json.Unmarshal(jsonFile, &data)
	if err != nil {
		panic(path + " Json file not found")
	}
	return data
}
