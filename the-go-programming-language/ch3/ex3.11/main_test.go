package main

import "testing"

func TestCommaFloatAndSign(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"one digit", "1", "1"},
		{"two digits", "12", "12"},
		{"three digits", "123", "123"},
		{"four digits", "1234", "1,234"},
		{"five digits", "12345", "12,345"},
		{"six digits", "123456", "123,456"},
		{"seven digits", "1234567", "1,234,567"},
		{"ten digits", "1234567890", "1,234,567,890"},
		{"negative four digits", "-1234", "-1,234"},
		{"positive sign four digits", "+1234", "+1,234"},
		{"decimal", "1.2345", "1.2345"},
		{"decimal four digit integer part", "1234.5678", "1,234.5678"},
		{"negative decimal", "-1234567.89", "-1,234,567.89"},
		{"negative decimal short", "-1234.5", "-1,234.5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CommaFloatAndSign(tt.in)
			if got != tt.want {
				t.Errorf("CommaFloatAndSign(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
