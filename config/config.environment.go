package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvironmentFileSetup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[Error]->Failed to load environment file : %s", err)
	}
}

func EnvironmentVariableCheck() {
	var requireEnvVariableList = []string{
		"TZ",
	}

	for _, v := range requireEnvVariableList {
		if os.Getenv(v) == "" {
			log.Fatalf("[Error]->Environment Variable ' %s ' is not set", v)
		}
	}
}