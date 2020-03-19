package api

import (
	"net/http"

	"github.com/go-rest/server"
	"github.com/gorilla/mux"
)

type Api struct {
	Srv       *server.Server
	Root      *mux.Router
	Coupons   *mux.Router
	Items     *mux.Router
	Employees *mux.Router
	Sales     *mux.Router
}

var API *Api

func InitRoutes(s *server.Server) {
	root := s.Handler.PathPrefix("/").Subrouter()

	API = &Api{
		Srv:  s,
		Root: root,
	}

	API.InitCouponRoutes()
	API.InitInventoryRoutes()
	API.InitEmployeeRoutes()
	API.InitSalesRoutes()
	API.Root.HandleFunc("/", MockHandler)
}

func MockHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("wololo"))
}
