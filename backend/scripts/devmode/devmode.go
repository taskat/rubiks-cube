package main

import (
	"devmode/restarter"
	"devmode/watcher"
	"fmt"
	"os"
	"regexp"
	"time"
)

func createRegex(s string) *regexp.Regexp {
	return regexp.MustCompile(s)
}

func createWatcher() *watcher.Watcher {
	r := restarter.NewRestarter()
	r.StartServer()
	folder := os.Args[1]
	includeString := os.Args[2]
	regexp := createRegex(includeString)
	return watcher.NewWatcher(folder, r, regexp)
}

func timer(tickChannel chan bool) {
	for {
		time.Sleep(500 * time.Millisecond)
		tickChannel <- true
	}
}

func readInput(inputChannel chan string) {
	for {
		var input string
		fmt.Scanln(&input)
		inputChannel <- input
	}
}

func main() {
	tickChannel := make(chan bool, 1)
	go timer(tickChannel)
	inputChannel := make(chan string, 1)
	go readInput(inputChannel)
	w := createWatcher()
	w.Watch(tickChannel, inputChannel)
}
