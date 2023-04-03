package restarter

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type Restarter struct {
	started bool
	pid     int
}

func NewRestarter() *Restarter {
	return &Restarter{}
}

func (r *Restarter) getPid() {
	fileName := "./rubik_server.pid"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic("Could not read pid file" + err.Error())
	}
	data = data[:len(data)-1]
	err = os.Remove(fileName)
	if err != nil {
		panic("Could not remove pid file" + err.Error())
	}
	r.pid, err = strconv.Atoi(string(data))
	if err != nil {
		panic("Could not parse pid" + err.Error())
	}
}

func (r *Restarter) killServer() {
	cmd := exec.Command("bash", "-c", "kill "+strconv.Itoa(r.pid))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		panic("Could not kill server" + err.Error())
	}
}

func (r *Restarter) removeServerExe() {
	err := os.Remove("./rubik_server.exe")
	if err != nil {
		_, isPathError := err.(*os.PathError)
		if !isPathError {
			panic("Could not remove server exe" + err.Error())
		}
	}
}

func (r *Restarter) RestartServer() {
	fmt.Println("Restarting server...")
	r.ShutDownServer()
	r.StartServer()
}

func (r *Restarter) ShutDownServer() {
	if r.started {
		r.killServer()
		r.waitForKill()
		r.removeServerExe()
	}
}

func (r *Restarter) StartServer() {
	cmd := exec.Command("bash", "-c", "./scripts/run_server_background.sh")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error starting server: " + err.Error())
	}
	r.started = err == nil
	if r.started {
		r.getPid()
	}
}

func (r *Restarter) waitForKill() {
	for {
		cmd := exec.Command("bash", "-c", "ps -p "+strconv.Itoa(r.pid))
		err := cmd.Run()
		if err != nil {
			_, isExitError := err.(*exec.ExitError)
			if !isExitError {
				panic("Could not check if server is running" + err.Error())
			}
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
