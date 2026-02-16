package main

import (
	"context"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"tutorial.sqlc.dev/app/tutorial"
)

func run() error {
	ctx := context.Background()
	conn, err := pgx.Connect(
		ctx,
		"host=localhost user=user password=password dbname=db sslmode=disable",
	)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := tutorial.New(conn)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio: pgtype.Text{
			String: "Co-author of the C Programming Language and the Go Programming Language",
			Valid:  true,
		},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	author, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// list all authors
	authors, err = queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, author))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
