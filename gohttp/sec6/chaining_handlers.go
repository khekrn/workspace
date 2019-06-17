package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func helloWebServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Http Server...............")
}

func helloGo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go !!!!!!!!!")
}

func helloJava(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Java !!!!!!!!!!!!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "localhost:9999",
	}

	http.HandleFunc("/", helloWebServer)
	http.HandleFunc("/go", log(helloGo))
	http.HandleFunc("/java", log(helloJava))

	server.ListenAndServe()
}
