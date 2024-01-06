package main

import (
	"log"
)

func main() {
	ginEngine := NewRouter()

	ginEngine.Run()
	log.Println("Server running...")
}
