package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/taskat/rubiks-cube/src/httpserver"
)

func main() {
	go func() {
		for {
			_, err := http.Get("http://localhost:8080/check")
			if err == nil {
				fmt.Println("Rubik's Server is running")
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	server := httpserver.NewServer()
	server.Start()
}
