package main

import "testing"

var sink int

func BenchmarkPopCountOriginal(b *testing.B) {
	for range b.N {
		sink = PopCount(0x123456789ABCDEF0)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for range b.N {
		sink = PopCountLoop(0x123456789ABCDEF0)
	}
}
