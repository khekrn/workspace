package main

import (
	"fmt"
	"net/http"
)

// MyHandler dummy handler
type MyHandler struct {
}

// ServeHTTP - handler function
func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello go !")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "localhost:9999",
		Handler: &handler,
	}

	server.ListenAndServe()
}
