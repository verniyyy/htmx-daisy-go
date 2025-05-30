package server

import (
	"encoding/json"
	"net/http"

	"github.com/verniyyy/htmx-daisy-go/internal/application/todo"
	"github.com/verniyyy/htmx-daisy-go/internal/view"
)

// NewMux creates a new HTTP ServeMux with a default route.
func NewMux() *http.ServeMux {
	r := NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(view.IndexHTML)
		if err != nil {
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
