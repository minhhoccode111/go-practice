package main

import (
	"context"
	"entdemo/ent"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open(
		"postgres",
		"host=localhost port=5432 user=user dbname=entdemo password=password sslmode=disable",
	)
	if err != nil {
		log.Fatalf("failed opening correction to postgres: %v", err)
	}
	defer client.Close()

	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
