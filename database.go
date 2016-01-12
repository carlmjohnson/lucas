package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const query = `
SELECT "spielberg_simpleredirect"."to_url",
       "spielberg_simpleredirect"."status_code"
FROM "spielberg_simpleredirect"
WHERE UPPER("spielberg_simpleredirect"."from_url"::text) = UPPER($1)
ORDER BY "spielberg_simpleredirect"."from_url"
ASC LIMIT 1;
`

var db *sql.DB

func init() {
	dbURL := os.Getenv("LUCAS_DATABASE_URL")
	var err error
	if db, err = sql.Open("postgres", dbURL); err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
