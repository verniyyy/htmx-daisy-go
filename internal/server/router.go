package server

import "net/http"

// Router is a custom HTTP router that wraps http.ServeMux.
type Router struct {
	*http.ServeMux
}

// NewRouter creates a new Router instance with a default ServeMux.
func NewRouter() *Router {
	return &Router{http.NewServeMux()}
}

// Get ...
func (r *Router) Get(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.route(http.MethodGet, pattern, handler)
}

// Head ...
func (r *Router) Head(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.route(http.MethodHead, pattern, handler)
}

// Post ...
func (r *Router) Post(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.route(http.MethodPost, pattern, handler)
}

// Put ...
func (r *Router) Put(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.route(http.MethodPut, pattern, handler)
}

// Patch ...
func (r *Router) Patch(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.route(http.MethodPatch, pattern, handler)
}

// Delete ...
func (r *Router) Delete(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.route(http.MethodDelete, pattern, handler)
}

// Connect ...
func (r *Router) Connect(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.route(http.MethodConnect, pattern, handler)
}

// Options ...
func (r *Router) Options(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.route(http.MethodOptions, pattern, handler)
}

// Trace ...
func (r *Router) Trace(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.route(http.MethodTrace, pattern, handler)
}

func (r *Router) route(method, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.ServeMux.HandleFunc(method+" "+pattern, handler)
}
