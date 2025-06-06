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

// set up the router in main
func main() {
	// create a router
	r := router.New()

	// register a hello route and its hello handler
	r.Handle("/hello", helloHandler)

	// simulate a request to /hello
	req := request.New("/hello", "GET", "", nil)

	// pass request to router
	res := r.Serve(req)

	// print response
	fmt.Printf("Status: %d\n", res.Status)
	fmt.Printf("Body: %s\n", res.Body)
}
