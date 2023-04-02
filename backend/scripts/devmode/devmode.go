package main

import (
	"devmode/restarter"
	"devmode/watcher"
	"os"
)

func main() {
	r := restarter.NewRestarter()
	r.StartServer()
	folder := os.Args[1]
	w := watcher.NewWatcher(folder, r.RestartServer)
	w.Watch()
}
