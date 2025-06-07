package main

import (
	"fmt"

	"github.com/dscholtz/baby-http.git/pkg/handler"
	"github.com/dscholtz/baby-http.git/pkg/request"
	"github.com/dscholtz/baby-http.git/pkg/router"
)

// create a simple hello handler
var helloHandler = handler.HandlerFunc(func(req *request.Request) *request.Response {
	return request.NewResponse(200, "Hello from baby-http!")
})

type AboutHandler struct{}

func (h AboutHandler) Serve(req *request.Request) *request.Response {
	return request.NewResponse(200, "This is a baby-http!!")
}

// set up the router in main
func main() {
	// create a router
	r := router.New()

	// register a hello route and its hello handler
	r.Handle("/hello", helloHandler)
	r.Handle("/about", AboutHandler{})

	// simulate a request to /hello
	req := request.New("/hello", "GET", "", nil)
	aboutReq := request.New("/about", "GET", "", nil)

	// pass request to router
	res := r.Serve(req)
	aboutRes := r.Serve(aboutReq)

	// print response
	fmt.Printf("Status: %d\n", res.Status)
	fmt.Printf("Body: %s\n", res.Body)
	fmt.Printf("Status: %d\n", aboutRes.Status)
	fmt.Printf("Body: %s\n", aboutRes.Body)
}
