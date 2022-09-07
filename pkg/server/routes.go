package server

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

func onShutdown(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Shutting server down\n")
}

func GetEndpoints() map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"/hello": getHello,
		"/":      getRoot,
	}
}
