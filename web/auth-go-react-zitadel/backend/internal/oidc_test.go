package internal

import "testing"

func TestHasAud(t *testing.T) {
	want := "urn:zitadel:iam:org:project:id:123:aud"
	cases := []struct {
		auds []string
		ok   bool
	}{
		{[]string{want}, true},
		{[]string{"other", want, "more"}, true},
		{[]string{"other"}, false},
		{nil, false},
	}
	for _, c := range cases {
		if got := hasAud(c.auds, want); got != c.ok {
			t.Errorf("hasAud(%v) = %v, want %v", c.auds, got, c.ok)
		}
	}
}
