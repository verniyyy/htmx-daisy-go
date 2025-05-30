package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// HTTPServer is a HTTP server.
type HTTPServer struct {
	srv           *http.Server
	shutdownHooks []ShutdownHook
}

// NewHTTPServer is a function to create a new HTTP server.
func NewHTTPServer(h http.Handler, host string, port int, shutdownHooks ...ShutdownHook) *HTTPServer {
	return &HTTPServer{
		srv: &http.Server{
			Addr:         net.JoinHostPort(host, strconv.Itoa(port)),
			BaseContext:  func(_ net.Listener) context.Context { return context.Background() },
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler:      h,
		},
		shutdownHooks: shutdownHooks,
	}
}

// Serve starts the HTTP server and listens for incoming requests.
func (s HTTPServer) Serve(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt)
	defer stop()

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- s.srv.ListenAndServe()
	}()

	select {
	case err := <-srvErr:
		return err
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	return s.shutdown(shutdownCtx)
}

func (s *HTTPServer) shutdown(ctx context.Context) error {

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	for _, hook := range s.shutdownHooks {
		hook(ctx)
	}

	return nil
}

// ShutdownHook is a function that is called when the server is shutting down.
type ShutdownHook func(ctx context.Context)
