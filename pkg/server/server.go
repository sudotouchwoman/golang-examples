package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func DefaultMux() *http.ServeMux {
	return http.NewServeMux()
}

type httpServer struct {
	server *http.Server
	mux    *http.ServeMux
}

func NewHttpServer(mux *http.ServeMux, addr string) httpServer {
	return httpServer{mux: mux, server: &http.Server{Addr: addr}}
}

func (srv *httpServer) RegisterEndpoints(routes map[string]func(http.ResponseWriter, *http.Request)) {
	if srv.mux == nil {
		log.Default().Println("HttpServer has no mux configured!")
		return
	}

	for r, fun := range routes {
		srv.mux.HandleFunc(r, fun)
	}
	log.Default().Printf("Registered %d new endpoints\n", len(routes))
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
	srv.mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		onShutdown(w, r)
		cancel()
	})

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
