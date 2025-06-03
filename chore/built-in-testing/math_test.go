package math

import "testing"

// Test Add function
func TestAdd(t *testing.T) {
	got := Add(2, 4)
	want := 6
	if got != want {
		t.Errorf("Add(2, 4) = %d; want %d", got, want)
	}
}
