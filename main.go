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
	services := &services.App{DAConfig: services.GetConfigurationDetails()}
	// s := gocron.NewScheduler(time.UTC)
	for _, scheduler := range services.DAConfig.Schedule {

		// s.Cron(scheduler.CronTime).Do(func() {
		services.Scheduler = scheduler
		validatedFiles := services.ValidateFiles()
		services.SendNotification(validatedFiles)
		// })

	}
	// s.StartBlocking()
}
