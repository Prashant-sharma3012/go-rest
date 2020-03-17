package server

import (
	"log"
	"net/http"

	"github.com/go-rest/api"
	"github.com/go-rest/db"
)

type Server struct {
	DB   *db.DB
	API  *api.Api
	port string
}

func Start() *Server {
	s := InitServer()

	log.Fatal(http.ListenAndServe(s.port, s.API.Handler))
	return s
}

func InitServer() *Server {
	return &Server{
		DB:   db.Connect(),
		API:  api.InitRoutes(),
		port: ":3000",
	}
}
