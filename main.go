package main

import (
	"github.com/joho/godotenv"
	"go-starter/config"
	"log"
)

func main() {
	// Set up dotenv
	err := godotenv.Load()

	config.Routes()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
