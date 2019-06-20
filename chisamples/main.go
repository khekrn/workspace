package main

import (
	"fmt"
	"net/http"

	"chisamples/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to nlp server...........")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Throttle(100))
	
	r.Get("/", defaultHandler)
	r.Post("/tokenize", handlers.TokenizeTextHandler)
	r.Mount("/debug", middleware.Profiler())
	http.ListenAndServe(":9999", r)
}
