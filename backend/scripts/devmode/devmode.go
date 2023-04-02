package main

import (
	"os"
)

func main() {
	restarter := newRestarter()
	restarter.startServer()
	folder := os.Args[1]
	watcher := newWatcher(folder, restarter.restartServer)
	watcher.watch()
}
