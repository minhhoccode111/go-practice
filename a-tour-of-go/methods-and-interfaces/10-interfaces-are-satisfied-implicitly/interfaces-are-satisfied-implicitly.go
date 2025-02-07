package main

import "fmt"

/*
Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no
explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its
implementation, which could then appear in any package without prearrangement.
*/

type IHuman interface {
	saySomething()
}
type Human struct {
	somethingToSay string
}

func (h Human) saySomething() {
	fmt.Println(h.somethingToSay)
}
func main() {
	var i IHuman = Human{"hello"}
	i.saySomething() // hello
}
