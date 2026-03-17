package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	var addr string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Name: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Print("Address: ")
	scanner.Scan()
	addr = scanner.Text()

	fmt.Printf("Name: %s - Address: %s\n", name, addr)

	obj := map[string]string{
		"name":    name,
		"address": addr,
	}

	data, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
