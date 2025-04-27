package main

import "fmt"

func main() {
	str := "Đặng Hoàng Minh"
	fmt.Printf("%T %[1]v\n", str)
	fmt.Printf("str length: %d\n", len(str))
	fmt.Printf("%T %[1]v\n", []rune(str))
	fmt.Printf("rune length: %d\n", len([]rune(str)))
	fmt.Printf("%T %[1]v\n", []byte(str))
	fmt.Printf("byte length: %d\n", len([]byte(str)))

	/*
	   $ go run main.go

	   string Đặng Hoàng Minh
	   str length: 19

	   []int32 [272 7863 110 103 32 72 111 224 110 103 32 77 105 110 104]
	   rune length: 15

	   []uint8 [196 144 225 186 183 110 103 32 72 111 195 160 110 103 32 77 105 110 104]
	   byte length: 19
	*/
}
