package persistent

import (
	"context"
	"errors"

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

func (r *UserRepo) Create(ctx context.Context, u *entity.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *UserRepo) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	var u entity.User
	err := r.db.WithContext(ctx).First(&u, id).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, entity.ErrNotFound
	case err != nil:
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var u entity.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, entity.ErrNotFound
	case err != nil:
		return nil, err
	}
	return &u, nil
}
