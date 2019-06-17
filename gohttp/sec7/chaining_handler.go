package main

import (
	"fmt"
	"net/http"
)

// RootHTTPHandler definition
type RootHTTPHandler struct{}

func (root *RootHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My name is root")
}

// GoHTTPHandler definition
type GoHTTPHandler struct{}

func (gohttp *GoHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My name is go !! ")
}

// JavaHTTPHandler definition
type JavaHTTPHandler struct{}

func (jhttp *JavaHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "My name is java !!")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "localhost:9999",
	}

	rh := new(RootHTTPHandler)
	gh := new(GoHTTPHandler)
	jh := new(JavaHTTPHandler)

	http.Handle("/", rh)
	http.Handle("/go", log(gh))
	http.Handle("/java", log(jh))

	server.ListenAndServe()
}
