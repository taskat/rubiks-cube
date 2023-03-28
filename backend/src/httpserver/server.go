package httpserver

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/taskat/rubiks-cube/src/config/handler"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s Server) getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func (s Server) configHandler(response http.ResponseWriter, request *http.Request) {
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	cube := confighandler.Handle("config.rubiks", string(content))
	if cube == nil {
		response.Write([]byte("ERROR"))
		printErrors()
	} else {
		data := cube.ToJSON()
		response.Write(data)
	}
}

func printErrors() {
	allMessages := eh.GetAllMessages()
	fmt.Println("======")
	for level, msgs := range allMessages {
		if len(msgs) == 0 {
			fmt.Println("There were no " + level + "s")
		} else {
			fmt.Println(level + "s:")
			for _, msg := range msgs {
				fmt.Println("  " + msg.String())
			}
		}
	}
}

func (s Server) Start() {
	http.HandleFunc("/config", s.configHandler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
