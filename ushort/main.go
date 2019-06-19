package main

import (
	"net/http"
	"ushort/handlers"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", handlers.Encode())
	mux.HandleFunc("/decode", handlers.Decode())

}
