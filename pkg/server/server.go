package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type httpServer struct {
	// Server that can be interrupted externally
	// using channels
	server *http.Server
	quit   chan bool
}

func New(m http.Handler, addr string) httpServer {
	return httpServer{
		server: &http.Server{
			Addr:    addr,
			Handler: m,
		},
		quit: make(chan bool, 1),
	}
}

func (srv *httpServer) Addr() string {
	return srv.server.Addr
}

func (srv *httpServer) Quit() chan<- bool {
	return srv.quit
}

func (srv *httpServer) Start() <-chan bool {
	log.Default().Printf("Starting serving at %s...", srv.Addr())
	// create channel to listen to system signals
	// exit peacefully once signal is recieved
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// launch server in a separate goroutine
	go func() {
		if err := srv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Default().Fatalf("Fatal in listen: %s\n", err)
		}
	}()

	// wait for the signal to appear
	// and inform the server to stop
	go func() {
		<-quit
		srv.quit <- true
	}()

	return srv.quit
}

func (srv *httpServer) ShutdownGracefully() {
	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// release resources/db connection here
		log.Default().Print("Releasing resources...")
		cancel()
	}()

	shutdownChan := make(chan error, 1)
	go func() { shutdownChan <- srv.server.Shutdown(timeout) }()

	select {
	case <-timeout.Done():
		log.Default().Fatal("Server shutdown timed out.")
	case err := <-shutdownChan:
		if err != nil {
			log.Default().Fatalf("Server shutdown aborted: %+v\n", err)
		} else {
			log.Default().Print("Server Exited.")
		}
	}
}
