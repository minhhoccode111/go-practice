package persistent

import (
	"context"
	"session-auth/entity"

	"gorm.io/gorm"
)

type SessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *SessionRepo {
	return &SessionRepo{
		db: db,
	}
}

func (r *SessionRepo) Create(ctx context.Context, s *entity.Session) error
func (r *SessionRepo) FindByTokenHash(ctx context.Context, hash string) (*entity.Session, error)
func (r *SessionRepo) DeleteByTokenHash(ctx context.Context, hash string) error
