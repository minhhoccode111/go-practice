package persistent

import (
	"context"

	"session-auth/entity"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, u *entity.User) error

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error)
