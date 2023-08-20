package main

import (
	"demo"
	"log"

	"github.com/google/uuid"
)

func main() {
	log.Println(demo.Message())
	log.Println(uuid.New().String())
	log.Println("Client demo")
	client := NewClient()
	client.get("http://example.com")
	log.Println("Server demo")
	server := NewServer()
	server.start()
}
