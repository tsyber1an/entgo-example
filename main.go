package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Funfun/entgo-example/ent"
	"github.com/Funfun/entgo-example/ent/book"
	playerEnt "github.com/Funfun/entgo-example/ent/player"
	_ "github.com/lib/pq"
)

func main() {
	log.Println("Welcome to Entgo example article")

	client, err := ent.Open("postgres", "user=postgres password=topsecret dbname=entgo_example sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// enable debugging
	client = client.Debug()

	log.Println("Connected to db")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Dump migration changes to stdout.
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	log.Println("press key <Enter> to continue")
	fmt.Scanln()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// create a Player
	player, err := client.Player.Create().
		SetNickname("John").
		SetEmail("info@tsyren.org").
		SetScores(1).
		Save(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("New player was created", player)

	player, err = client.Player.Query().Where(playerEnt.Nickname("John")).Only(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("Found player was created", player)

	leo, err := client.Person.Create().SetName("Leo Tolstoy").Save(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("create a person record", leo)

	publishedAt, _ := time.Parse(time.RFC3339, "1869-01-01T15:04:06Z07:00")
	warAndPeace, err := client.Book.Create().SetTitle("War and Peace").SetAuthor(leo).SetCreatedAt(publishedAt).Save(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("create a book", warAndPeace)

	// query
	res, err := client.Book.Query().Where(book.Title("War and Peace")).Only(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("query book's author", res)

	author, err := client.Book.Query().Where(book.Title("War and Peace")).QueryAuthor().Only(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Println("query book's author", author)
}
