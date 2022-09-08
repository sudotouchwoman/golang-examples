package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sudotouchwoman/golang-examples/pkg/api"
	"github.com/sudotouchwoman/golang-examples/pkg/server"
)

func main() {

	var host, port string
	flag.StringVar(&host, "h", "localhost", "Server Host.")
	flag.StringVar(&port, "p", "2000", "Server Port.")
	flag.Parse()

	router := mux.NewRouter()

	// separate subrouters for public and private sections
	// let one apply middleware for specific routes
	// e.g., policy-related stuff
	public_router := router.PathPrefix("/public").Subrouter()
	private_router := router.PathPrefix("/private").Subrouter()

	// once all the handlers are registered, define
	// 404 handler (otherwise, this one acts like a wildcard
	// and bypasses all the other middlewares)
	router.Use(server.PanicRecovery, server.LogResponseCode)
	router.NotFoundHandler = router.
		NewRoute().
		HandlerFunc(http.NotFound).
		GetHandler()

	private_router.Use(server.LogRemoteAddr)

	// add public/private endpoints
	quit := make(chan bool, 1)
	api.AddEndpoints(private_router, api.ShutdownEndpoint(api.Shutdown, quit))
	api.AddEndpoints(public_router, api.ApiPublicEndpoints()...)
	api.AddEndpoints(router, api.HealthEndpoint())

	srv := server.New(router, host+":"+port)

	go func() {
		// propagate signal to the server
		// haha arrows go brr
		srv.Quit() <- <-quit
	}()

	<-srv.Start()
	srv.ShutdownGracefully()
}
