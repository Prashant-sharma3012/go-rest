package main

import (
	"fmt"

	"github.com/go-rest/api"
	"github.com/go-rest/server"
)

func main() {
	fmt.Println("Initialize server")
	srv := server.InitServer()
	fmt.Println("Initialize routes")
	api.InitRoutes(srv)

	fmt.Println("All DOne")
	srv.Start()

	fmt.Println("server is up and running")
}
