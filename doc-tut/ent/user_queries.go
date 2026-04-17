package main

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/user"
	"fmt"
	"log"
)

func CreateUser(ctx context.Context, client *ent.Client, age int, name string) (*ent.User, error) {
	u, err := client.User.Create().SetAge(age).SetName(name).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client, name string) (*ent.User, error) {
	u, err := client.User.Query().Where(user.Name(name)).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
