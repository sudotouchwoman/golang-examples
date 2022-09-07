package server

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestHttpServer_MuProperty(t *testing.T) {
	expectedAddr := "localhost:8080"
	mu := mux.NewRouter()
	srv := NewHttpServer(mu, expectedAddr)

	if !(mu == srv.mu && mu == srv.server.Handler) {
		t.Errorf("Pointers to Router do not match!")
	}

	got := srv.Addr()
	if got != expectedAddr {
		t.Errorf("srv.Addr() = %s; want %s", got, expectedAddr)
	}
}
