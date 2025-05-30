package server

import (
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/verniyyy/htmx-daisy-go/assets"
	"github.com/verniyyy/htmx-daisy-go/internal/application/todo"
	"github.com/verniyyy/htmx-daisy-go/internal/view"
)

// NewMux creates a new HTTP ServeMux with a default route.
func NewMux() *http.ServeMux {
	r := NewRouter()

	r.Get("/static/", http.StripPrefix("/static", http.FileServer(http.FS(assets.Assets))).ServeHTTP)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// w.WriteHeader(http.StatusOK)

		u := todo.NewListUseCase()
		todos, err := u.Execute()
		if err != nil {
			slog.ErrorContext(r.Context(), "Failed to execute template", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.New("index.templ").ParseFS(view.IndexTemplate, "index.templ")
		if err != nil {
			slog.ErrorContext(r.Context(), "Failed to execute template", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		args := map[string]any{"todos": todos}
		if err := tmpl.Execute(w, args); err != nil {
			slog.ErrorContext(r.Context(), "Failed to execute template", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	r.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		u := todo.NewListUseCase()
		todos, err := u.Execute()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		serveJSON(w, http.StatusOK, todos)
	})

	return r.ServeMux
}

func serveJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// WithLog ...
func WithLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), r.Method, slog.String("url", r.URL.String()))
		h.ServeHTTP(w, r)
	})
}
