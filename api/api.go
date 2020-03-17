package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Api struct {
	Handler *mux.Router
}

func InitRoutes() *Api {
	r := mux.NewRouter()
	r.HandleFunc("/", MockHandler).Methods("Get")

	return &Api{
		Handler: r,
	}
}

func MockHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("wololo"))
}
