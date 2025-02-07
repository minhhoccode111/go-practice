package main

import "fmt"

/*
Exercise: Stringers
Make the IPAddr type implement fmt.Stringer to print the address as a dotted
quad.

For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
*/

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
		// loopback: 127.0.0.1
		// googleDNS: 8.8.8.8
	}
}
