package main

import "fmt"

func main() {
	fmt.Println(commaFloatAndSign("1"))
	fmt.Println(commaFloatAndSign("12"))
	fmt.Println(commaFloatAndSign("123"))
	fmt.Println(commaFloatAndSign("1234"))
	fmt.Println(commaFloatAndSign("12345"))
	fmt.Println(commaFloatAndSign("123456"))
	fmt.Println(commaFloatAndSign("1234567"))
	fmt.Println(commaFloatAndSign("12345678"))
	fmt.Println(commaFloatAndSign("123456789"))
	fmt.Println(commaFloatAndSign("1234567890"))
}

func commaFloatAndSign(s string) string {
	return ""
}
