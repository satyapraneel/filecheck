package main

import (
	"filecheck/services"

	"github.com/joho/godotenv"
)

func main() {
	schedule()
}

func schedule() {

	if err := godotenv.Load("env.yaml"); err != nil {
		panic(err)
	}
	services := &services.App{DAStruct: services.GetConfigurationDetails()}
	validatedFiles := services.ValidateFiles()
	services.SendNotification(validatedFiles)
}
