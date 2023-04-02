package watcher

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

type Watcher struct {
	lastEdited   map[string]time.Time
	folder       string
	callback     func()
	startup      bool
	includeFiles *regexp.Regexp
}

func NewWatcher(folder string, callback func(), includeFiles *regexp.Regexp) *Watcher {
	w := &Watcher{folder: folder, callback: callback, startup: true, includeFiles: includeFiles}
	w.lastEdited = make(map[string]time.Time)
	w.updateFiles(w.folder)
	w.updateLastEdited()
	w.startup = false
	return w
}

func (w *Watcher) log(msg string) {
	if !w.startup {
		fmt.Println(msg)
	}
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
			if !w.includeFiles.MatchString(fileName) {
				continue
			}
			if _, ok := w.lastEdited[fileName]; !ok {
				w.log("Added file: " + fileName)
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
			w.log("Deleted file: " + file)
			delete(w.lastEdited, file)
			updated = true
			continue
		}
		lastMod := fileInfo.ModTime()
		if lastMod != modTime {
			w.log("Updated file: " + file)
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
