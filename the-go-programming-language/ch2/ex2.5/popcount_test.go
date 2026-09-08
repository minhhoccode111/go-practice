package main

import "testing"

var sink int

func BenchmarkPopCountOriginal(b *testing.B) {
	for range b.N {
		sink = PopCount(0x123456789ABCDEF0)
	}
}

func BenchmarkPopCountKernighan(b *testing.B) {
	for range b.N {
		sink = PopCountKernighan(0x123456789ABCDEF0)
	}
}
