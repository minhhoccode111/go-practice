package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"session-auth/entity"
	"session-auth/helpers"
	"session-auth/repo"
)

type Service struct {
	u repo.UserRepo
	s repo.SessionRepo
}

func New(
	userRepo repo.UserRepo,
	sessionRepo repo.SessionRepo,
) *Service {
	return &Service{
		u: userRepo,
		s: sessionRepo,
	}
}

func (s *Service) Register(ctx context.Context, email, password string) (*entity.User, error) {
	email = strings.ToLower(email)
	if !helpers.IsValidEmail(email) {
		return nil, entity.ErrInvalidEmail
	}
	if helpers.IsValidPassword(password) {
		return nil, entity.ErrWeakPassword
	}
	if _, err := s.u.FindByEmail(ctx, email); err == nil {
		return nil, entity.ErrEmailTaken
	} else if !errors.Is(err, entity.ErrNotFound) {
		return nil, fmt.Errorf("find user: %w", err)
	}
	hash, err := helpers.GenerateHash(password)
	if err != nil {
		return nil, fmt.Errorf("hash password: %w", err)
	}
	u := &entity.User{Email: email, PasswordHash: hash}
	if err := s.u.Create(ctx, u); err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return u, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (token string, err error) {
	panic("")
}

func (s *Service) Authenticate(ctx context.Context, token string) (*entity.User, error) {
	panic("")
}

func (s *Service) Logout(ctx context.Context, token string) error {
	panic("")
}
