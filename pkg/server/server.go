package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func onServerClosed() {
	fmt.Println("Server stopped serving.")
}

func onServerError(err error) {
	fmt.Printf("Encountered error: %s\n", err)
	os.Exit(1)
}

func GetMux() *http.ServeMux {
	return http.NewServeMux()
}

func RunServer(port string, mux *http.ServeMux) {
	// starts serving on given port
	// and blocks until some error occurs
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	err := http.ListenAndServe(port, mux)

	if errors.Is(err, http.ErrServerClosed) {
		onServerClosed()
	} else if err != nil {
		onServerError(err)
	}
}
