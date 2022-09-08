package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote := quotesAPI.fetchQuote()
	io.WriteString(w, quote.String())
}

func getManyQuotes(w http.ResponseWriter, r *http.Request) {
	nquotes := mux.Vars(r)["nquotes"]
	num, err := strconv.Atoi(nquotes)
	if err != nil {
		log.Default().Printf("Failed to convert %s to int", nquotes)
		panic(err)
	}
	if num <= 0 {
		log.Default().Printf("Invalid number of quotes: %d", num)
		panic(num)
	}

	quotes := quotesAPI.fetchQuotes(uint(num))
	if err := json.NewEncoder(w).Encode(quotes); err != nil {
		panic(err)
	}
}
