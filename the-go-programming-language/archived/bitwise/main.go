package main

import "fmt"

/*
&  : bitwise AND
|  : bitwise OR
^  : bitwise XOR
&^ : bit clear (AND NOT)
<< : left shift
>> : right shift
*/

func p(x int) {
	fmt.Printf("%08b\n", x)
}

func main() {
	x := 1<<1 | 1<<5
	y := 1<<1 | 1<<2
	p(x)      // the set {1, 5}
	p(y)      // the set {1, 2}
	p(x | y)  // the set {1, 2, 5}
	p(x & y)  // the set {1}
	p(x ^ y)  // the symmetric difference {2, 5}
	p(x &^ y) // the difference {5} (clear y from x)

	for i := range uint(8) {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	p(x << 1) // the set {2, 6}
	p(y << 2) // the set {3, 4}

	fmt.Printf("%v\n", 0xdeadbeef)
}
