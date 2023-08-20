package main

import (
	"demo"
	"log"

	"github.com/google/uuid"
)

func main() {
	log.Println(demo.Message())
	log.Println(uuid.New().String())
	server := NewServer()
	server.start()
}
