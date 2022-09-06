package main

import (
	"fmt"

	"github.com/sudotouchwoman/golang-examples/pkg/server"
)

func main() {
	fmt.Println("Starting serving...")

	mux := server.GetMux()
	server.HandleRoutes(mux, server.GetRoutes())
	server.RunServer(":2000", mux)
}
