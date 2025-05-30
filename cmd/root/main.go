package main

import (
	"context"
	"log/slog"

	"github.com/verniyyy/htmx-daisy-go/internal/server"
)

func main() {
	mux := server.WithLog(server.NewMux())
	srv := server.NewHTTPServer(mux, "", 8080, func(ctx context.Context) {
		slog.InfoContext(ctx, "Server is shutting down gracefully")
	})

	if err := srv.Serve(context.Background()); err != nil {
		panic(err)
	}
}
