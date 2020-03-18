package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-rest/db"
	"github.com/gorilla/mux"
)

type Server struct {
	DB      *db.DB
	Handler *mux.Router
	port    string
}

func (s *Server) Start() *Server {
	log.Fatal(http.ListenAndServe(s.port, s.Handler))
	return s
}

func InitServer() *Server {
	fmt.Println("Connecting to DB")
	fmt.Println("Creating HTTP handler")

	return &Server{
		DB:      db.Connect(),
		Handler: mux.NewRouter(),
		port:    ":3000",
	}
}
