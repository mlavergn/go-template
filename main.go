package main

import (
	"demo"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println(demo.Message())
	fmt.Println(uuid.New().String())
}
