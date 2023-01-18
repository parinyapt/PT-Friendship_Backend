package config

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvironmentSetup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[Error]->Failed to load environment file : %s", err)
	}
}