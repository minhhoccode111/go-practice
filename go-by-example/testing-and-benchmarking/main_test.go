package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestMinBasic(t *testing.T) {
	ans := IntMin(1, 2)
	if ans != 1 {
		t.Errorf("IntMin(1, 2) = %d; want ", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 1},
		{-1, 2, -1},
		{01, 02, 01},
		{-0, 0, 0},
		{3, 2, 2},
		{4, 2, 2},
		{-100, 2, -100},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		IntMin(1, 2)
	}
}
