package router

import (
	"github.com/dscholtz/baby-http.git/pkg/handler"
	"github.com/dscholtz/baby-http.git/pkg/request"
)

type Router struct {
	routes map[string]handler.Handler // path -> handler
}

// constructor
func New() *Router {
	return &Router{
		routes: make(map[string]handler.Handler),
	}
}

// Handle() registers a path to a handler during setup.
func (r *Router) Handle(path string, h handler.Handler) {
	r.routes[path] = h
}

func (r *Router) Serve(req *request.Request) *request.Response {
	// Send in the path to the router, if there is a handler corresponding to that path
	// return that handler and serve the request to it.
	h, ok := r.routes[req.Path]
	if !ok {
		return request.NewResponse(404, "Not found")
	}
	return h.Serve(req)
}
