package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var pf = fmt.Printf

func main() {
	p := point{1, 2}

	pf("struct1: %v\n", p)
	pf("struct2: %+v\n", p)
	pf("struct3: %#v\n", p)
	pf("type   : %T\n", p)
	pf("boolean: %t\n", true)
	pf("int    : %d\n", 123)
	pf("bin    : %b\n", 123)
	pf("char   : %c\n", 123)
	pf("hex    : %x\n", 123)
	pf("float1 : %f\n", 123.123)
	pf("float2 : %e\n", 123.123)
	pf("float3 : %E\n", 123.123)
	pf("str1   : %s\n", "123")
	pf("str2   : %q\n", "123")
	pf("str3   : %x\n", "123")
	pf("pointer: %p\n", &p)
	pf("width1 : |%6d|%6d|\n", 12, 345)
	pf("width2 : |%6.2f|%6.2f|\n", 1.2, 3.45)
	pf("width3 : |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	pf("width4 : |%6s|%6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io     : an %s\n", "error")

	/*
	   $ go run string-formatting.go
	   struct1: {1 2}
	   struct2: {x:1 y:2}
	   struct3: main.point{x:1, y:2}
	   type   : main.point
	   boolean: true
	   int    : 123
	   bin    : 1111011
	   char   : {
	   hex    : 7b
	   float1 : 123.123000
	   float2 : 1.231230e+02
	   float3 : 1.231230E+02
	   str1   : 123
	   str2   : "123"
	   str3   : 313233
	   pointer: 0xc000010130
	   width1 : |    12|   345|
	   width2 : |  1.20|  3.45|
	   width3 : |1.20  |3.45  |
	   width4 : |   foo|     b|
	   sprintf: a string
	   io     : an error
	*/
}
