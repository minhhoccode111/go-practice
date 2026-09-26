package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func IsValidPassword(password string) bool {
	return len(password) < 8
}

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error generate password: %w", err)
	}
	return string(hash), nil
}
