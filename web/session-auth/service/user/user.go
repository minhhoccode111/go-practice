package user

import (
	"context"
	"session-auth/entity"
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

func (s *Service) Register(ctx context.Context, email, password string) (*entity.User, error)
func (s *Service) Login(ctx context.Context, email, password string) (token string, err error)
func (s *Service) Authenticate(ctx context.Context, token string) (*entity.User, error)
func (s *Service) Logout(ctx context.Context, token string) error
