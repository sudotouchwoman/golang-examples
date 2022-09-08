package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	// stores information about endpoints
	// for a RESTful API
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

func handleEndpoint(mux *mux.Router, r Route) {
	mux.
		Methods(r.Method).
		Path(r.Pattern).
		Name(r.Name).
		Handler(r.Handler)
}

func AddEndpoints(mux *mux.Router, routes ...Route) *mux.Router {
	for _, r := range routes {
		handleEndpoint(mux, r)
	}
	return mux
}

func ShutdownEndpoint(handler http.HandlerFunc, quit chan<- bool) Route {
	// wraps the given shutdown handler
	// to populate a channel once request to the /shutdown endpoint
	// is handled
	return Route{
		"Shutdown",
		http.MethodGet,
		"/shutdown",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.ServeHTTP(w, r)
			quit <- true
		}),
	}
}

func HealthEndpoint() Route {
	return Route{
		"HealthCheck",
		http.MethodGet,
		"/health",
		http.HandlerFunc(getHealthCheck),
	}
}

func ApiPublicEndpoints() []Route {
	// returns slice of public endpoints
	// to be consumed by RegisterEndpoints
	return []Route{
		{
			"Index",
			http.MethodGet,
			"/",
			http.HandlerFunc(getRoot),
		},
		{
			"Hello",
			http.MethodGet,
			"/hello",
			http.HandlerFunc(getHello),
		},
	}
}
