package service

import (
	"context"

	"session-auth/entity"
)

type (
	User interface {
		Register(ctx context.Context, email, password string) (*entity.User, error)
		Login(ctx context.Context, email, password string) (token string, err error)
		Authenticate(ctx context.Context, token string) (*entity.User, error)
		Logout(ctx context.Context, token string) error
	}
)
