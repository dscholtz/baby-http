package handler

import "github.com/dscholtz/baby-http.git/pkg/request"

// Any type that has a Serve(*Request) method that returns a *Response satisfies this.
// Think of it like: "anything that can handle a request."
type Handler interface {
	Serve(req *request.Request) *request.Response
}

// HandlerFunc is a type definition for a function that:
// - takes a *Request
// - returns a *Response
//
// This allows you to treat regular functions as first-class handler values.
type HandlerFunc func(*request.Request) *request.Response

// This method gives HandlerFunc the Serve() method it needs to
// satisfy the Handler interface.
//
// So if you write a plain function that matches HandlerFunc's shape,
// you can wrap it like this:
//
//	handler := HandlerFunc(myFunc)
//
// and now handler.Serve(...) will work!
func (f HandlerFunc) Serve(req *request.Request) *request.Response {
	return f(req)
}
