package main

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/car"
	"entdemo/ent/group"
	"entdemo/ent/user"
	"log"
)

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// first, create the users
	mhc, err := client.User.Create().SetAge(25).SetName("minhhoccode111").Save(ctx)
	if err != nil {
		return err
	}

	neta, err := client.User.Create().SetAge(28).SetName("Neta").Save(ctx)
	if err != nil {
		return err
	}

	// then create the cars and attach them to the users created above
	err = client.Car.Create().SetModel("Tesla").SetOwner(mhc).Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.Create().SetModel("Mazda").SetOwner(mhc).Exec(ctx)
	if err != nil {
		return err
	}
	// create the groups and add their users in the creation
	err = client.Group.Create().SetName("GitLab").AddUsers(mhc, neta).Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Group.Create().SetName("GitHub").AddUsers(mhc).Exec(ctx)
	if err != nil {
		return err
	}
	log.Println("the graph was created successfully")
	return nil
}

// get all user's cars within the gorup named 'github'
func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.Query().
		Where(group.Name("GitHub")). // (Group(Name=Github),)
		QueryUsers().                // (User(Name=Ariel,age=30))
		QueryCars().
		// (Car(Model=Tesla, RegisteredAt=<time>), Car(Model=Mazda, RegisteredAt=<time>),)
		All(ctx)
	if err != nil {
		return err
	}
	log.Println("cars returned: ", cars)
	return nil
}

func QueryMHCCars(ctx context.Context, client *ent.Client) error {
	// get 'mhc' from previous steps
	mhc := client.User.Query().Where(user.HasCars(), user.Name("mhc")).OnlyX(ctx)
	cars, err := mhc.QueryGroups(). // get the group that mhc is connected to:
					QueryUsers(). // (Group(Name=GitHub), Group(Name=GitLab))
					QueryCars().  //
					Where(car.Not(car.Model("Mazda"))).
					All(ctx)

	if err != nil {
		return err
	}
	log.Printf("cars returned: ", cars)
	// output: (Car(Model=Tesla,RegisteredAt=<time>), Car(Model=Ford,RegisteredAt=<time>),)
	return nil
}

// get all groups that have users (query with a look-aside predicate)
func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	groups, err := client.Group.Query().Where(group.HasUsers()).All(ctx)
	if err != nil {
		return err
	}
	log.Println("groups returned:", groups)
	// output: (Group(Name=GitHub), Group(Name=GitLab), )
	return nil
}
