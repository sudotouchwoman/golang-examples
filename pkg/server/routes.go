package server

import (
	"io"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, _ *http.Request) {
	log.Default().Println("got / requrest")
	io.WriteString(w, "You made it to the root!\n")
}

func getHello(w http.ResponseWriter, _ *http.Request) {
	log.Default().Println("got /hello request")
	io.WriteString(w, "Hello there!\n")
}

func GetRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"hello": getHello,
		"root":  getRoot,
	}
}

func HandleRoutes(mux *http.ServeMux, routes map[string]func(http.ResponseWriter, *http.Request)) {
	for r, fun := range routes {
		mux.HandleFunc(r, fun)
	}
}
