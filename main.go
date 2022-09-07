package main

import (
	"flag"
	"log"

	"github.com/gorilla/mux"
	"github.com/sudotouchwoman/golang-examples/pkg/server"
)

func main() {

	var host, port string
	flag.StringVar(&host, "h", "localhost", "Server Host.")
	flag.StringVar(&port, "p", "2000", "Server Port.")
	flag.Parse()

	srv := server.NewHttpServer(mux.NewRouter(), host+":"+port)
	// apply some middleware to intercept
	srv.WrapHandler(server.LogRequests)

	// add endpoint handlers
	srv.RegisterEndpoints(server.GetEndpoints(), "GET")

	log.Default().Printf("Starting serving at %s", srv.Addr())
	// starts serving on given port
	// and blocks until some error occurs
	srv.ServeGracefully()

	log.Default().Printf("Server Exited.")
}
