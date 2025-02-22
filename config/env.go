package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return
	}
	log.Println("Config loaded")
}
