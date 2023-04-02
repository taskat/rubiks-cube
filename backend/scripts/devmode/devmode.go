package main

import (
	"os"
	"os/exec"
	"strconv"
	"time"
)

func getPid() int {
	fileName := "./rubik_server.pid"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	data = data[:len(data)-1]
	err = os.Remove(fileName)
	if err != nil {
		panic(err)
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		panic(err)
	}
	return pid
}

func startServer() int {
	cmd := exec.Command("bash", "-c", "./scripts/run_server_background.sh")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return getPid()
}

func killServer(pid int) {
	cmd := exec.Command("bash", "-c", "kill "+strconv.Itoa(pid))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
}

func removeServerExe() {
	err := os.Remove("./rubik_server.exe")
	if err != nil {
		_, isPathError := err.(*os.PathError)
		if !isPathError {
			panic(err)
		}
	}
}

func waitForKill(pid int) {
	for {
		cmd := exec.Command("bash", "-c", "ps -p "+strconv.Itoa(pid))
		err := cmd.Run()
		if err != nil {
			_, isExitError := err.(*exec.ExitError)
			if !isExitError {
				panic(err)
			}
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func restartServer(pid int) int {
	killServer(pid)
	waitForKill(pid)
	removeServerExe()
	return startServer()
}

func main() {
	pid := startServer()
	folder := os.Args[1]
	watcher := newWatcher(folder, restartServer, pid)
	watcher.watch()
}
