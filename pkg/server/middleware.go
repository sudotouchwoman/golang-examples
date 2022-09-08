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

func LogRemoteAddr(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("%s %v (%s)", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func LogResponseCode(next http.Handler) http.Handler {
	// log incoming requests and their corresponding response codes
	// the response code is 200 by default as custom handlers may not
	// call the WriteHeader method directly
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxyWriter := &proxyResponseWriter{ResponseWriter: w, code: 200}
		next.ServeHTTP(proxyWriter, r)
		log.Default().Printf("%s %v - %d", r.Method, r.URL, proxyWriter.code)
	})
}

func PanicRecovery(next http.Handler) http.Handler {
	// recover if for some reason execution panics
	// and respond with a 500 status code
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Default().Printf("Panicked: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
