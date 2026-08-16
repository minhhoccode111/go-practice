package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x0 float64 = 3.4
	v0 := reflect.ValueOf(x0)
	fmt.Println("type:", v0.Type())
	fmt.Println("kind is float64:", v0.Kind() == reflect.Float64)
	fmt.Println("value:", v0.Float())

	fmt.Println()

	var x1 uint8 = 'x'
	v1 := reflect.ValueOf(x1)
	fmt.Println("type:", v1.Type())                           // uint8
	fmt.Println("kind is uint8:", v1.Kind() == reflect.Uint8) // true
	x1 = uint8(v1.Uint())                                     // v1.Uint returns a uint64

	fmt.Println()

	type MyInt int
	var x2 MyInt = 7
	v2 := reflect.ValueOf(x2)
	fmt.Println("value:", v2)

	fmt.Println()

	var x3 float64 = 3.4
	v3 := reflect.ValueOf(x3)
	fmt.Println("settability of v3:", v3.CanSet())

	fmt.Println()

	var x4 float64 = 3.4
	p4 := reflect.ValueOf(&x4)                     // note: take the address of x4
	fmt.Println("type of p4:", p4.Type())          // *float64
	fmt.Println("settability of p4:", p4.CanSet()) // false
	v4 := p4.Elem()
	fmt.Println("settability of v4:", v4.CanSet()) // true
	v4.SetFloat(7.1)
	fmt.Println(v4.Interface()) // 7.1
	fmt.Println(x4)             // 7.1

	fmt.Println()

	type T5 struct {
		A int
		B string
	}
	t5 := T5{23, "skidoo"}
	s5 := reflect.ValueOf(&t5).Elem()
	typeOfT := s5.Type()
	for i := range s5.NumField() {
		f5 := s5.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f5.Type(), f5.Interface())
	}
	s5.Field(0).SetInt(77)
	s5.Field(1).SetString("sunset strip")
	fmt.Println("t5 is now", t5)
}
