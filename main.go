package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Welcome to Entgo example article")

	db, err := sql.Open("postgres", "user=postgres password=topsecret dbname=entgo_example sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to db")
}
