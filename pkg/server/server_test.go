package server

import (
	"net/http"
	"testing"
)

func TestNewHttpServer(t *testing.T) {
	expectedAddr := "localhost:8080"
	srv := New(http.NewServeMux(), expectedAddr)

	if got := srv.Addr(); got != expectedAddr {
		t.Errorf("srv.Addr() = %s; want %s", got, expectedAddr)
	}

	expectedEmpty := 0
	if got := len(srv.Quit()); got != expectedEmpty {
		t.Errorf("len(srv.Quit()) = %d; want %d", got, expectedEmpty)
	}
}
