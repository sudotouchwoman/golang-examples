package server

import (
	"log"
	"net/http"
)

type proxyResponseWriter struct {
	// helper struct to extract http response code
	// once handler returns
	http.ResponseWriter
	code int
}

func (rwi *proxyResponseWriter) WriteHeader(code int) {
	rwi.code = code
	rwi.ResponseWriter.WriteHeader(code)
}

func LogRequests(handler http.Handler) http.Handler {
	// log incoming requests and their corresponding response codes
	// the response code is 200 by default as custom handlers may not
	// call the WriteHeader method directly
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxyWriter := &proxyResponseWriter{ResponseWriter: w, code: 200}
		handler.ServeHTTP(proxyWriter, r)
		log.Default().Printf("%s %v - %d", r.Method, r.URL, proxyWriter.code)
	})
}
