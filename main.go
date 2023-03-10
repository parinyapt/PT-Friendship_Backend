package main

import (
	"os"

	"github.com/parinyapt/PT-Friendship_Backend/config"
	"github.com/parinyapt/PT-Friendship_Backend/database"
	"github.com/parinyapt/PT-Friendship_Backend/routes"
)

func main() {
	config.FlagSetup()
	if os.Getenv("DEPLOY_MODE") == "development" {
		config.EnvironmentFileSetup()
	}
	if os.Getenv("DEPLOY_MODE") == "production" {
		config.GinReleaseModeSetup()
	}
	config.EnvironmentVariableCheck()
	config.TimezoneSetup()

	database.Connect()
	
	routes.Setup()
}