package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := range ticker.C {
		fmt.Println(i)
	}
}
