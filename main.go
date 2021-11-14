package main

import (
	"log"

	"github.com/Funfun/entgo-example/ent"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Welcome to Entgo example article")

	client, err := ent.Open("postgres", "user=postgres password=topsecret dbname=entgo_example sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	log.Println("Connected to db")
}
