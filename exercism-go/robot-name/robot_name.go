package robotname

import (
	"errors"
	"math/rand/v2"
	"strconv"
)

var used = map[string]bool{}

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	maxName := 26 * 26 * 10 * 10 * 10

	for {
		if len(used) == maxName {
			return "", errors.New("Maximum name reached")
		}

		name := randomName()
		if r.validName(name) {
			r.name = name
			used[name] = true
			break
		}
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func (r *Robot) validName(name string) bool {
	_, ok := used[name]
	return !ok
}

func randomName() (name string) {
	name += randomLetter()
	name += randomLetter()
	name += randomDigit()
	name += randomDigit()
	name += randomDigit()
	return
}

func randomLetter() string {
	n := rune('A' + rand.IntN(26))
	return string(n)
}

func randomDigit() string {
	n := strconv.Itoa(rand.IntN(10))
	return n
}
