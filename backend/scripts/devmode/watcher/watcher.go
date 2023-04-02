package watcher

import (
	"fmt"
	"os"
	"time"
)

type Watcher struct {
	lastEdited map[string]time.Time
	folder     string
	callback   func()
}

func NewWatcher(folder string, callback func()) *Watcher {
	w := &Watcher{folder: folder}
	w.lastEdited = make(map[string]time.Time)
	w.updateFiles(w.folder)
	w.updateLastEdited()
	w.callback = callback
	return w
}

func (w *Watcher) updateFiles(currentFolder string) {
	entries, err := os.ReadDir(currentFolder)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			w.updateFiles(currentFolder + "/" + entry.Name())
		} else {
			fileName := currentFolder + "/" + entry.Name()
			if fileName[len(fileName)-4:] == ".exe" {
				continue
			}
			if _, ok := w.lastEdited[fileName]; !ok {
				fmt.Println("Added file: " + fileName)
				w.lastEdited[fileName] = time.Time{}
			}
		}
	}
}

func (w *Watcher) updateLastEdited() bool {
	updated := false
	for file, modTime := range w.lastEdited {
		fileInfo, err := os.Lstat(file)
		if err != nil {
			fmt.Println("Deleted file: " + file)
			delete(w.lastEdited, file)
			updated = true
			continue
		}
		lastMod := fileInfo.ModTime()
		if lastMod != modTime {
			fmt.Println("Updated file: " + file)
			w.lastEdited[file] = lastMod
			updated = true
		}
	}
	return updated
}

func (w *Watcher) Watch() {
	for {
		time.Sleep(500 * time.Millisecond)
		w.updateFiles(w.folder)
		if w.updateLastEdited() {
			w.callback()
		}
	}
}
