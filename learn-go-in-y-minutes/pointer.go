package main

import (
	"fmt"
	"unsafe"
)

func main() {
	exampleFullyGarbageCollected()

	exampleValidPointerUse()

	exampleInvalidPointerUse()
}

// function returns a pointer
func createObject() *int {
	x := 42   // x is allocated on the heap (Go decides this automatically)
	return &x // return a pointer to x
}

// go is fully garbage collected, it has pointers but no pointer arithmetic
func exampleFullyGarbageCollected() {
	p := createObject()
	// even though x's scope ended. Go's garbage collector ensures p is valid until it's no longer referenced
	fmt.Println(*p) // 42
	// when p is no longer used, the garbage collector frees the memory
}

// go allows pointers to references memory addresses, but you cannot perform
// arithmetic (e.g. ptr++, ptr + offset) to navigate memory. this prevent unsafe
// memory access
func exampleValidPointerUse() {
	a := 10
	var ptr *int = &a // pointer to a
	fmt.Println(*ptr) // 10 (dereferencing works)
	fmt.Println(ptr)  // view address
	*ptr = 20         // modify a through the pointer
	fmt.Println(a)    // 20
}

func exampleInvalidPointerUse() {
	arr := [3]int{1, 2, 3}
	ptr := &arr[0]
	// ptr++ // compile error "cannot convert pointer to unsafe.Pointer"
	fmt.Println(ptr)
}

/*
WHY THIS MATTERS

safety: no pointer arithmetic prevents buffer overfolows, dangling pointers and
other memory corruption bugs

simplicity: garbage collection reduces manual memory management errors (e.g.
memory leaks)
*/

// for rare cases requiring pointer arithmetic (e.g. interacting with hardware),
// go provides the `unsafe` package, but it's discouraged for typical use
func exampleUnsafePointer() {
	arr := [3]int{1, 2, 3}
	ptr := uintptr(unsafe.Pointer(&arr[0]))
	ptr += unsafe.Sizeof(arr[0]) // risky and not idiomatic
	// convert back to pointer: unsafe.Pointer(ptr)
}
