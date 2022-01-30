package main

import (
	"filecheck/services"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("env.yaml"); err != nil {
		panic(err)
	}
	services := &services.App{DAStruct: services.GetConfigurationDetails()}
	validatedFiles := services.ValidateFiles()
	services.SendNotification(validatedFiles)
}
