package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	datasource := flag.String("pgDs", "postgres://postgres:0000@postgres:5432/meepshop?sslmode=disable", "postgres datasource")
	flag.Parse()

	db, err := sql.Open("postgres", *datasource)
	if err != nil {
		log.Fatalf("pg connect err: %s", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("pg ping err: %s", err)
	}
	log.Printf("pg ping OK")
}
