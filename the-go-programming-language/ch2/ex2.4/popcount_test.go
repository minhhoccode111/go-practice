package main

import "testing"

var sink int

func BenchmarkPopCountOriginal(b *testing.B) {
	for range b.N {
		sink = PopCount(0x123456789ABCDEF0)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for range b.N {
		sink = PopCountShift(0x123456789ABCDEF0)
	}
}
