package cipher

import (
	"regexp"
	"strings"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance > 25 || distance < -25 || distance == 0 {
		return nil
	}
	return shift{distance: distance}
}

func (c shift) Encode(input string) string {
	input = regexp.MustCompile(`[^a-z]`).ReplaceAllString(strings.ToLower(input), "")
	result := []rune{}
	for _, v := range input {
		v = (v+26-'a'+rune(c.distance))%26 + 'a'
		result = append(result, v)
	}
	return string(result)
}

func (c shift) Decode(input string) string {
	// opposite to encode
	opposite := shift{distance: -c.distance}
	return opposite.Encode(input)
}

func NewVigenere(key string) Cipher {
	re := regexp.MustCompile(`^[a-z]+$`)
	if !re.MatchString(key) || regexp.MustCompile(`^a+$`).MatchString(key) {
		return nil
	}
	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	input = regexp.MustCompile(`[^a-z]+`).ReplaceAllString(strings.ToLower(input), "")
	result := ""
	for i, r := range input {
		result += shift{distance: int(v.key[i%len(v.key)]) - 'a'}.Encode(string(r))
	}
	return result
}

func (v vigenere) Decode(input string) string {
	input = regexp.MustCompile(`[^a-z]+`).ReplaceAllString(strings.ToLower(input), "")
	result := ""
	for i, r := range input {
		result += shift{distance: int(v.key[i%len(v.key)]) - 'a'}.Decode(string(r))
	}
	return result
}
