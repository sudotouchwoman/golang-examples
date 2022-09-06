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

func onShutdown(w http.ResponseWriter, _ *http.Request) {
	log.Default().Println("got /shutdown request")
	io.WriteString(w, "Shutting server down\n")
}

func GetRoutes() map[string]func(http.ResponseWriter, *http.Request) {
	return map[string]func(http.ResponseWriter, *http.Request){
		"/hello": getHello,
		"/":      getRoot,
	}
}
