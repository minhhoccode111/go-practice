package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
	fmt.Println(comma("12345678"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1234567890"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	buf.Grow(n + n%3)
	first := n % 3
	if first == 0 {
		first = 3
	}
	buf.WriteString(s[:first])
	for i := first; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
