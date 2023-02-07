package main

import (
	"dependencies/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Get env variables
	err := godotenv.Load("../config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Generate all routes
	r := router.Generate()

	r.Run()
}
