package main

import (
	"database/sql"
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

	row := db.QueryRow(query, r.URL.Path)

	var (
		to_url      string
		status_code int
	)

	err := row.Scan(&to_url, &status_code)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		log.Fatalf("FATAL DATABASE TROUBLE: %v", err)
	}

	w.Header().Set("Location", to_url)
	w.WriteHeader(status_code)
}

func main() {
	flag.Parse()

	log.Printf("Starting lucas on %s...\n\n", *ListenOn)

	if err := http.ListenAndServe(*ListenOn, http.HandlerFunc(redirector)); err != nil {
		log.Fatalf("FATAL ERROR: %s\n", err)
	}
}
