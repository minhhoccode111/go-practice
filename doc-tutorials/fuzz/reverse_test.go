package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{
		"Hello, World!",
		" ",
		"!12345",
	}

	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, revErr := Reverse(orig)
		if revErr != nil {
			t.Skip()
		}
		doubleRev, doubleRevErr := Reverse(rev)
		if doubleRevErr != nil {
			t.Skip()
		}
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d\n", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q\n", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q\n", rev)
		}
	})
}
