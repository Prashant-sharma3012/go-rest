package main

import (
	"fmt"

	"github.com/go-rest/server"
)

func main() {
	fmt.Println("starting server")
	server.Start()
	fmt.Println("server is up and running")
}
