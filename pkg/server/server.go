package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

type httpServer struct {
	// keeps pointer to the original mux object,
	// which gives opportunity to directly call its methods
	// i.e., to update its handlers
	server *http.Server
	mu     *mux.Router
}

func NewHttpServer(m *mux.Router, addr string) httpServer {
	return httpServer{mu: m, server: &http.Server{Addr: addr, Handler: m}}
}

func (srv *httpServer) Addr() string {
	return srv.server.Addr
}

func (srv *httpServer) RegisterEndpoints(routes map[string]func(http.ResponseWriter, *http.Request), httpmethods ...string) {
	if srv.mu == nil {
		log.Default().Println("HttpServer has no mux configured!")
		return
	}

	for r, fun := range routes {
		srv.mu.HandleFunc(r, fun).Methods(httpmethods...)
	}
}

func (srv *httpServer) WrapHandler(wrapper ...func(http.Handler) http.Handler) {
	// applies given sequence to the server's handler
	// note that the order of passed wrappers is preserved
	for _, w := range wrapper {
		srv.server.Handler = w(srv.server.Handler)
	}
}

func (srv *httpServer) ServeGracefully() {
	// create channel to listen to system signals
	// exit peacefully when signal is recieved
	gotSignal := make(chan os.Signal, 1)
	signal.Notify(gotSignal, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// stop the server on shutdown endpoint
	// this may actually imply some validation steps
	srv.mu.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		onShutdown(w, r)
		defer cancel()
	}).Methods(http.MethodGet)

	// launch server in a separate goroutine
	go func() {
		if err := srv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Default().Fatalf("Fatal in listen: %s\n", err)
		}
	}()

	// block this goroutine until given signal occurs
	// or the context is released by request to /shutdown
	select {
	case <-gotSignal:
	case <-ctx.Done():
		if err := srv.server.Shutdown(ctx); err != nil {
			log.Default().Fatalf("Server shutdown aborted: %+v\n", err)
		}
	}
}
