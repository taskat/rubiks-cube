package main

import (
	"fmt"

	"github.com/taskat/rubiks-cube/src/httpserver"
)

func main() {
	fmt.Println("Starting server...")
	server := httpserver.NewServer()
	server.Start()
}
