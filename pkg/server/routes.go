package server

import (
	"io"
	"log"
	"net/http"
)

func LogRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("Got a %s request for: %v", r.Method, r.URL)
		handler.ServeHTTP(w, r)
		// At this stage, our handler has "handled" the request
		// but we can still write to the client there
		// but we won't do that
		log.Default().Println("Handler finished processing request")
	})
}

func getRoot(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "You made it to the root!\n")
}

func getHello(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello there!\n")
}

func onShutdown(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Shutting server down\n")
}

func GetEndpoints() map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"/hello": getHello,
		"/":      getRoot,
	}
}
