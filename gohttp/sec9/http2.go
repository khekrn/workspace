package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
)

// MyHandler struct definition
type MyHandler struct{}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Go HTTP2 Server")
}

func main() {
	myHandler := new(MyHandler)
	server := http.Server{
		Addr:    "localhost:9999",
		Handler: myHandler,
	}

	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServe()
}
