package entity

import "errors"

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrWeakPassword = errors.New("weak password")
	ErrEmailTaken   = errors.New("email already registered")
	ErrNotFound     = errors.New("not found")

	ErrInvalidCredentials = errors.New("invalid credentials")
)
