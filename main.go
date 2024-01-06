package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("couldn't load environment file")
	}

	ginEngine := NewRouter()

	ginEngine.Run()
	log.Println("Server running...")
}
