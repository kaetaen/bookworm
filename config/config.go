package config

import (
	"log"

	"github.com/joho/godotenv"
)

var Env map[string]string
var err error

func LoadEnv() {
	Env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
