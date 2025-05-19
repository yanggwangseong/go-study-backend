package main

import (
	"log"
)

func main() {
	server := NewServer()
	log.Println("Server starting on :8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
