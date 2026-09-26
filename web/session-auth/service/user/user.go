package user

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"session-auth/entity"
	"session-auth/helpers"
	"session-auth/repo"
)

const sessionTTL = 7 * 24 * time.Hour

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
	email = strings.ToLower(email)
	u, err := s.u.FindByEmail(ctx, email)
	switch {
	case errors.Is(err, entity.ErrNotFound):
		return "", entity.ErrInvalidCredentials
	case err != nil:
		return "", fmt.Errorf("find user: %w", err)
	}
	if err := helpers.VerifyHash(u.PasswordHash, password); err != nil {
		return "", entity.ErrInvalidCredentials
	}
	token, err = helpers.GenerateToken()
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}
	session := &entity.Session{
		UserID:    u.ID,
		TokenHash: helpers.HashToken(token),
		ExpiresAt: time.Now().Add(sessionTTL),
	}
	if err := s.s.Create(ctx, session); err != nil {
		return "", fmt.Errorf("create session: %w", err)
	}
	return token, nil
}

func (s *Service) Authenticate(ctx context.Context, token string) (*entity.User, *entity.Session, error) {
	session, err := s.s.FindByTokenHash(ctx, helpers.HashToken(token))
	switch {
	case errors.Is(err, entity.ErrNotFound):
		return nil, nil, entity.ErrInvalidCredentials
	case err != nil:
		return nil, nil, fmt.Errorf("find session: %w", err)
	}
	if time.Now().After(session.ExpiresAt) {
		return nil, nil, entity.ErrInvalidCredentials
	}
	u, err := s.u.FindByID(ctx, session.UserID)
	if err != nil {
		return nil, nil, fmt.Errorf("find user: %w", err)
	}
	return u, session, nil
}

func (s *Service) ListSessions(ctx context.Context, userID uint) ([]entity.Session, error) {
	sessions, err := s.s.ListByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("list sessions: %w", err)
	}
	return sessions, nil
}

func (s *Service) DeleteSession(ctx context.Context, userID, sessionID uint) error {
	if err := s.s.DeleteByIDAndUserID(ctx, sessionID, userID); err != nil {
		return fmt.Errorf("delete session: %w", err)
	}
	return nil
}

func (s *Service) DeleteAllSessions(ctx context.Context, userID uint) error {
	if err := s.s.DeleteByUserID(ctx, userID); err != nil {
		return fmt.Errorf("delete sessions: %w", err)
	}
	return nil
}

func (s *Service) Logout(ctx context.Context, token string) error {
	if token == "" {
		return nil
	}
	if err := s.s.DeleteByTokenHash(ctx, helpers.HashToken(token)); err != nil {
		return fmt.Errorf("delete session: %w", err)
	}
	return nil
}
