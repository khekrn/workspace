package main

import "net/http"

func main() {
	server := http.Server{
		Addr:    "localhost:9999",
		Handler: nil,
	}
	server.ListenAndServe()
}
