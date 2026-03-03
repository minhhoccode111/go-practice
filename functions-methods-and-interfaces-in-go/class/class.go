package main

import "fmt"

type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}

func main() {
	v := MyInt(5)
	fmt.Println(v.Double())
}
