package main

import (
	"flag"
	"log"

	"net/http"
)

var (
	Debug    = flag.Bool("debug", true, "Log activity")
	ListenOn = flag.String("listen-on", "127.0.0.1:8080", "Address:port to listen on")
)

func redirector(w http.ResponseWriter, r *http.Request) {
	if *Debug {
		log.Printf("Redirecting %s ...\n", r.URL.Path)
	}

	urlStr := "http://example.com"
	code := 302

	w.Header().Set("Location", urlStr)
	w.WriteHeader(code)
}

func main() {
	flag.Parse()

	log.Printf("Starting lucas on %s...\n\n", *ListenOn)

	if err := http.ListenAndServe(*ListenOn, http.HandlerFunc(redirector)); err != nil {
		log.Fatalf("FATAL ERROR: %s\n", err)
	}
}
