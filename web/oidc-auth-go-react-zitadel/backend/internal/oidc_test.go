package internal

import "testing"

func TestHasAud(t *testing.T) {
	// setup.sh writes the raw numeric project id (not the urn scope form) to
	// backend/.env, and that is what ZITADEL puts in the access token aud.
	want := "385060034775613444"
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
