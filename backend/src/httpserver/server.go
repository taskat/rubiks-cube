package httpserver

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/taskat/rubiks-cube/src/config/handler"
	eh "github.com/taskat/rubiks-cube/src/errorhandler"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()
	s := &Server{mux: mux}
	s.mux.HandleFunc("/config", s.configHandler)
	return s
}

func (s Server) addCORSHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.WriteHeader(http.StatusOK)
		return
	}
	s.mux.ServeHTTP(w, r)
}

func (s Server) getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func (s Server) configHandler(response http.ResponseWriter, request *http.Request) {
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	errorHandler := eh.NewHandler()
	cube := confighandler.Handle("config.rubiks", string(content), errorHandler)
	if cube == nil {
		response.Write([]byte("ERROR"))
		errorHandler.PrintAll()
	} else {
		fmt.Println("Successful parse")
		data := cube.ToJSON()
		response.Write(data)
	}
}

func (s Server) Start() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(s.addCORSHeaders))
	fmt.Println(err)
}
