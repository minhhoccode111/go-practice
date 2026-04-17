package main

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/car"
	"fmt"
	"log"
)

func CreateCars(
	ctx context.Context,
	client *ent.Client,
	carNames []string,
	userName string,
	userAge int,
) (*ent.User, error) {
	carCreates := make([]*ent.CarCreate, 0, len(carNames))
	// loop through each car name to create new car
	for _, v := range carNames {
		carCreates = append(carCreates, client.Car.Create().SetModel(v))
	}

	// then bulk insert
	cars, err := client.Car.CreateBulk(carCreates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating cars: %w", err)
	}

	// create a new user, and add it the created cars
	u, err := client.User.Create().SetAge(userAge).SetName(userName).AddCars(cars...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryCars(ctx context.Context, u *ent.User) error {
	cars, err := u.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars: ", cars)

	// what about filtering specific cars.
	ford, err := u.QueryCars().Where(car.Model("Ford")).Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user car: %w", err)
	}
	log.Println("returned cars: ", ford)
	return nil
}

func QueryCarUsers(ctx context.Context, u *ent.User) error {
	cars, err := u.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying cars: %w", err)
	}

	// query the inverse edge
	for _, c := range cars {
		owner, err := c.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", c, err)
		}
		log.Printf("car %q owner: %q\n", c, owner)
	}
	return nil
}
