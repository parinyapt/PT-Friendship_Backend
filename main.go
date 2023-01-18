package main

import (
	"os"

	"github.com/parinyapt/PT-Friendship_Backend/config"
	"github.com/parinyapt/PT-Friendship_Backend/database"
)

func main() {
	config.FlagSetup()

	if os.Getenv("DEPLOY_MODE") == "development" {
		config.EnvironmentFileSetup()
	}
	config.EnvironmentVariableCheck()
	database.Connect()
	config.TimezoneSetup()
}