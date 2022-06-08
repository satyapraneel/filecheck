package main

import (
	"filecheck/services"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	schedule()
}

func schedule() {

	env := "env.yaml"
	if len(os.Args) >= 2 {
		env = os.Args[1]
	}
	if err := godotenv.Load(env); err != nil {
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
