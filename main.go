package main

import (
	"flag"
	"log"

	"github.com/sudotouchwoman/golang-examples/pkg/server"
)

func main() {

	var host, port string
	flag.StringVar(&host, "h", "localhost", "Server Host. Defaults to localhost")
	flag.StringVar(&port, "p", "2000", "Server Port. Defaults to 2000")
	flag.Parse()

	addr := host + ":" + port

	srv := server.NewHttpServer(server.DefaultMux(), addr)
	srv.RegisterEndpoints(server.GetRoutes())

	log.Default().Printf("Starting serving at %s", srv.Addr())
	// starts serving on given port
	// and blocks until some error occurs
	srv.ServeGracefully()

	log.Default().Printf("Server Exited.")
}
