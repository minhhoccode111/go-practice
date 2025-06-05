package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := `
	{
		"name": "minhhoccode111",
		"age": 18
	}
	`
	var user User
	if err := json.Unmarshal([]byte(jsonData), &user); err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	fmt.Printf("User: %+v\n", user)
}
