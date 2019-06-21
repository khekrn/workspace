package main

import (
	"flag"
	"fmt"
	"net/http"

	"chisamples/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to nlp server...........")
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "enable debug log")

	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if *debug{
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Throttle(100))

	r.Get("/", defaultHandler)
	r.Post("/tokenize", handlers.TokenizeTextHandler)
	r.Post("/wc", handlers.WordCountHandler)
	r.Mount("/debug", middleware.Profiler())
	http.ListenAndServe(":9999", r)
}
