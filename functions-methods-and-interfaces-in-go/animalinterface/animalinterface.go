package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (Cow) Eat()   { fmt.Println("grass") }
func (Cow) Move()  { fmt.Println("walk") }
func (Cow) Speak() { fmt.Println("moo") }

func (Bird) Eat()   { fmt.Println("worms") }
func (Bird) Move()  { fmt.Println("fly") }
func (Bird) Speak() { fmt.Println("peep") }

func (Snake) Eat()   { fmt.Println("mice") }
func (Snake) Move()  { fmt.Println("slither") }
func (Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := make(map[string]Animal)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		parts := strings.Fields(scanner.Text())
		if len(parts) != 3 {
			continue
		}

		cmd := parts[0]
		name := parts[1]
		arg := parts[2]

		switch cmd {

		case "newanimal":
			var a Animal

			switch arg {
			case "cow":
				a = Cow{}
			case "bird":
				a = Bird{}
			case "snake":
				a = Snake{}
			default:
				fmt.Println("Unknown animal")
				continue
			}

			animals[name] = a
			fmt.Println("Created it!")

		case "query":
			a, ok := animals[name]
			if !ok {
				fmt.Println("Animal not found")
				continue
			}

			switch arg {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			}
		default:
			fmt.Println("Usage:")
			fmt.Println("newanimal [name] [cow|bird|snake]")
			fmt.Println("query [name] [eat|move|speak]")
		}

	}
}
