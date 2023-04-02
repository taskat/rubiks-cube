package main

import (
	"devmode/restarter"
	"devmode/watcher"
	"os"
	"regexp"
)

func createRegex(s string) *regexp.Regexp {
	return regexp.MustCompile(s)
}

func main() {
	r := restarter.NewRestarter()
	r.StartServer()
	folder := os.Args[1]
	includeString := os.Args[2]
	w := watcher.NewWatcher(folder, r.RestartServer, createRegex(includeString))
	w.Watch()
}
