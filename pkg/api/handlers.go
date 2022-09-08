package api

import (
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "You made it to the root!\n")
}

func getHello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello there!\n")
}

func Shutdown(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Shutting server down\n")
}

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	// borrowed from github.com/gorilla/mux docs
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
