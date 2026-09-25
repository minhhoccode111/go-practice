package helpers

import "net/mail"

func IsValidEmail(s string) bool {
	addr, err := mail.ParseAddress(s)
	return err == nil && addr.Address == s
}
