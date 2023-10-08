package httpserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/taskat/rubiks-cube/src/algo/handler"
	"github.com/taskat/rubiks-cube/src/config/handler"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
	"github.com/taskat/rubiks-cube/src/executor"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()
	s := &Server{mux: mux}
	s.mux.HandleFunc("/", s.check)
	s.mux.HandleFunc("/config", s.configHandler)
	s.mux.HandleFunc("/all", s.allHandler)
	return s
}

func (s Server) addHeaders(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	response.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	response.Header().Set("Content-Type", "application/json")
	if request.Method == "OPTIONS" {
		response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		response.Header().Set("Access-Control-Max-Age", "86400")
		response.WriteHeader(http.StatusOK)
		return
	}
	s.mux.ServeHTTP(response, request)
}

func (s Server) allHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		response.Header().Set("Allow", "POST")
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		s.writeError(response, http.StatusBadRequest, err)
		return
	}
	var content Request
	err = json.Unmarshal(body, &content)
	if err != nil {
		s.writeError(response, http.StatusBadRequest, err)
		return
	}
	errorHandler := eh.NewHandler()
	cube := confighandler.Handle("config.rubiks", string(content.Config), &errorHandler)
	if len(errorHandler.GetErrors()) != 0 || len(errorHandler.GetWarnings()) != 0 {
		msg := "There are errors/warnings in the configuration. Skipping algorithm checking and execution."
		errorHandler.AddInfo(eh.NewContext(0, 0), msg, "algorithm.algo")
	}
	algo := algohandler.Handle("algorithm.algo", string(content.Algo), &errorHandler, cube)
	executor := executor.NewExecutor(&errorHandler)
	turns, steps := executor.Execute(cube.Clone(), algo)
	fmt.Println("Turns: ", turns)
	fmt.Println("Steps: ", steps)
	messages := errorHandler.GetMessages()
	fmt.Println("Messages: ", messages)
	result := NewResult(cube, messages, turns, steps)
	data, err := json.Marshal(result)
	if err != nil {
		s.writeError(response, http.StatusInternalServerError, err)
		return
	}
	response.Write(data)
}

func (s Server) check(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Rubik's Cube server is running!"))
}

func (s Server) configHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		response.Header().Set("Allow", "POST")
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		s.writeError(response, http.StatusBadRequest, err)
		return
	}
	var content Request
	err = json.Unmarshal(body, &content)
	if err != nil {
		s.writeError(response, http.StatusBadRequest, err)
		return
	}
	errorHandler := eh.NewHandler()
	cube := confighandler.Handle("config.rubiks", string(content.Config), &errorHandler)
	messages := errorHandler.GetMessages()
	fmt.Println("Messages: ", messages)
	result := NewResult(cube, messages, map[string][]string{}, []string{})
	data, err := json.Marshal(result)
	if err != nil {
		s.writeError(response, http.StatusInternalServerError, err)
		return
	}
	response.Write(data)
}

func (s Server) Start() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(s.addHeaders))
	fmt.Println(err)
}

func (s Server) writeError(response http.ResponseWriter, code int, err error) {
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(err.Error()))
}
