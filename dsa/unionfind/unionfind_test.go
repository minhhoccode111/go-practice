package unionfind

import (
	"math/rand"
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	n := 10
	uf := New(10)

	if uf.Count() != n {
		t.Fatalf("expected count %d, got %d", n, uf.Count())
	}

	for i := range n {
		if !uf.Connected(i, i) {
			t.Fatalf("element %d should be connected to itself", i)
		}
	}

	for i := range n {
		for j := i + 1; j < n; j++ {
			if uf.Connected(i, j) {
				t.Fatalf("element %d and %d should not be connected", i, j)
			}
		}
	}
}

func TestSingleUnion(t *testing.T) {
	uf := New(5)

	uf.Union(1, 2)

	if !uf.Connected(1, 2) {
		t.Fatal("1 and 2 should be connected after union")
	}

	if uf.Count() != 4 {
		t.Fatalf("expected 4 sets, but got %d", uf.Count())
	}
}

func TestTransitiveUnion(t *testing.T) {
	uf := New(5)
	uf.Union(1, 2)
	uf.Union(2, 3)

	if !uf.Connected(1, 3) {
		t.Fatal("transitive violated: 1 and 3 should be connected")
	}
}

func TestUnionIdempotent(t *testing.T) {
	uf := New(5)

	uf.Union(1, 2)
	before := uf.Count()

	uf.Union(1, 2)
	after := uf.Count()

	if before != after {
		t.Fatal("union should not reduce count when already connected")
	}
}

func TestSelfUnion(t *testing.T) {
	uf := New(5)
	before := uf.Count()
	uf.Union(3, 3)
	after := uf.Count()
	if before != after {
		t.Fatal("self union should not change set count")
	}
}

func TestMultipleComponents(t *testing.T) {
	uf := New(6)
	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(4, 5)

	if uf.Count() != 3 {
		t.Fatalf("expected 3 components, got %d", uf.Count())
	}

	uf.Union(1, 3)
	if !uf.Connected(0, 2) {
		t.Fatal("0 and 2 should be connected after merge")
	}

	if uf.Count() != 2 {
		t.Fatalf("expected 2 components, got %d", uf.Count())
	}
}

func TestFindConsistency(t *testing.T) {
	uf := New(5)
	uf.Union(0, 1)
	uf.Union(1, 2)
	root := uf.Find(0)
	for _, v := range []int{1, 2} {
		if uf.Find(v) != root {
			t.Fatalf("node %d has inconsistent root", v)
		}
	}
}

func TestRandomized(t *testing.T) {
	n := 100
	uf := New(n)

	ref := make([]int, n)
	for i := range ref {
		ref[i] = i
	}

	// naive union
	unionRef := func(a, b int) {
		ra, rb := ref[a], ref[b]
		for i := range ref {
			if ref[i] == rb {
				ref[i] = ra
			}
		}
	}

	connectedRef := func(a, b int) bool {
		return ref[a] == ref[b]
	}

	for range 500 {
		a := rand.Intn(n)
		b := rand.Intn(n)

		uf.Union(a, b)
		unionRef(a, b)

		for x := range n {
			for y := range n {
				if uf.Connected(x, y) != connectedRef(x, y) {
					t.Fatalf("mismatch at (%d,%d)", x, y)
				}
			}
		}
	}
}

func TestInvalidIndex(t *testing.T) {
	uf := New(5)
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for invalid index")
		}
	}()

	uf.Find(100)
}

func TestLargeInput(t *testing.T) {
	n := 100_000
	uf := New(n)
	for i := 1; i < n; i++ {
		uf.Union(0, i)
	}

	if uf.Count() != 1 {
		t.Fatal("all nodes should be connected")
	}
}

func FuzzUnionFind(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		n := 50
		if a < 0 || b < 0 {
			return
		}
		a %= n
		b %= n

		uf := New(n)
		uf.Union(a, b)

		if !uf.Connected(a, b) {
			t.Fatal("should be connected")
		}
	})
}
