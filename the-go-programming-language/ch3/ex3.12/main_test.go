package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{"anagram", "listen", "silent", true},
		{"single char", "a", "a", true},
		{"both empty", "", "", true},
		{"not anagram", "hello", "world", false},
		{"different lengths", "abc", "ab", false},
		{"repeated letters", "aabb", "abab", true},
		{"same letters different count", "aab", "abb", false},
		{"identical", "abc", "abc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.s1, tt.s2); got != tt.want {
				t.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
