package main

import (
	"fmt"
	"net/http"
)

func helloGo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello go !!!!!")
}

func helloJava(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Java !!!!!")
}

func helloWebServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Http WebServer")
}

func main() {
	server := http.Server{
		Addr: "localhost:9999",
	}

	http.HandleFunc("/", helloWebServer)
	http.HandleFunc("/java", helloJava)
	http.HandleFunc("/go", helloGo)

	server.ListenAndServe()
}
