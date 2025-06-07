package middleware

import (
	"fmt"

	"github.com/dscholtz/baby-http.git/pkg/handler"
	"github.com/dscholtz/baby-http.git/pkg/request"
)

// Logging takes handler as an argument. Then it wraps this handler
// with extra behaviour. (logging in this instance)
func Logging(h handler.Handler) handler.Handler {
	return handler.HandlerFunc(func(req *request.Request) *request.Response {
		fmt.Printf("[LOG] %s %s\n", req.Method, req.Path)
		return h.Serve(req)
	})
}
