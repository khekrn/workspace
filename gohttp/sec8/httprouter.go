package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func root(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "Hello root !!!")
}

func greet(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "Hello "+p.ByName("lang"))
}

func greetV2(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	req.ParseForm()
	fmt.Fprintln(w, "Hello "+req.FormValue("lang"))
}

func log(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		fmt.Printf("Calling handler - %T\n", h)
		h(w, req, p)
	}
}

func main() {
	mux := httprouter.New()
	mux.GET("/", root)
	mux.GET("/greet/:lang", log(greet))
	mux.GET("/greetv2/", log(greetV2))
	server := http.Server{
		Addr:    "localhost:9999",
		Handler: mux,
	}

	server.ListenAndServe()

}
