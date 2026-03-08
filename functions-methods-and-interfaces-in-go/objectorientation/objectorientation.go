package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat()   { fmt.Println(a.food) }
func (a *Animal) Move()  { fmt.Println(a.locomotion) }
func (a *Animal) Speak() { fmt.Println(a.noise) }

func printUsage() {
	fmt.Println("Usage: > [cow|bird|snake] [eat|move|speak]")
}

func printInvalid() {
	fmt.Println("invalid input")
}

func main() {
	printUsage()
	fmt.Println("Example input : > cow eat")
	fmt.Println("Example output: grass")

	dict := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			line := scanner.Text()
			paths := strings.Fields(strings.TrimSpace(line))
			if len(paths) != 2 {
				printInvalid()
				printUsage()
				continue
			}
			a, ok := dict[paths[0]]
			if !ok {
				printInvalid()
				printUsage()
				continue
			}
			switch paths[1] {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				printInvalid()
				printUsage()
			}
		}
	}
}
