package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type User struct {
	ID           int       `db:"id"`
	Username     string    `db:"username"`
	Password     string    `db:"-"` // should be included
	ApiKey       string    `db:"-"` // should be included
	privateField string    `db:"private_field"`
	CreatedAt    time.Time `db:"created_at"`
}

func main() {
	u1 := User{
		ID:           1,
		Username:     "minhhoccode111",
		Password:     "supersecretpassword",
		privateField: "private field",
		CreatedAt:    time.Now(),
	}

	s1, err := parse(u1)
	if err != nil {
		panic(err)
	}

	s2, err := parse(&u1)
	if err != nil {
		panic(err)
	}

	// both struct and pointer struct work
	fmt.Println(s1) // ["Password", "ApiKey"]
	fmt.Println(s2) // ["Password", "ApiKey"]
}

func parse(s any) ([]string, error) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New("item is not a struct")
	}
	result := make([]string, 0)
	t := v.Type()
	fields := t.Fields()
	for v := range fields {
		tag := v.Tag.Get("db")
		if v.IsExported() && tag == "-" {
			result = append(result, v.Name)
		}
	}
	return result, nil
}
