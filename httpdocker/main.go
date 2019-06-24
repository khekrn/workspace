package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// MyServer Demo web server
type MyServer struct{}

func (ms *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	greetingHandler(w, r)
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("in greetingHandler")
	fmt.Fprintln(w, "go http server......................")
	log.Info("return from greetingHandler")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("in getHandler")

	defer r.Body.Close()
	r.ParseForm()

	id := r.FormValue("id")
	name := r.FormValue("name")

	fmt.Fprintln(w, "Id - ", id, ", Name - ", name)

	log.Info("return from getHandler")
}

func awesome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "awesome..............................")
}

func main() {
	fmt.Println("hello go.....")
	log.SetFormatter(&log.JSONFormatter{})

	// server := http.Server{
	// 	Addr:         "localhost:9999",
	// 	Handler:      nil,
	// 	ReadTimeout:  10 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	http.HandleFunc("/", greetingHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/awesome", awesome)

	//server.ListenAndServe()

	http.ListenAndServe(":9999", nil)
}
