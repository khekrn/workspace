package main

import (
	"fmt"
	"net/http"
)

// GreetingHandler hello world handler
type GreetingHandler struct{}

func (gh *GreetingHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Hello Http Server !!!!")
}

// GoHandler golang handler
type GoHandler struct{}

func (goh *GoHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "I love go these days !!!!")
}

// JavaHandler java handler
type JavaHandler struct{}

func (jh *JavaHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "these days java is bit boring......... :(-")
}

func main() {

	greeting := GreetingHandler{}
	golang := GoHandler{}
	jhandler := JavaHandler{}

	server := http.Server{
		Addr:    "localhost:9999",
		Handler: nil,
	}

	http.Handle("/", &greeting)
	http.Handle("/java", &jhandler)
	http.Handle("/go", &golang)

	server.ListenAndServe()

}
