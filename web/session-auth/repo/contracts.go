package repo

import (
	"context"

	"session-auth/entity"
)

type (
	UserRepo interface {
		Create(context.Context, *entity.User) error
		FindByEmail(context.Context, string) (*entity.User, error)
		FindByID(context.Context, uint) (*entity.User, error)
	}
	SessionRepo interface {
		Create(context.Context, *entity.Session) error
		FindByTokenHash(context.Context, string) (*entity.Session, error)
		ListByUserID(context.Context, uint) ([]entity.Session, error)
		DeleteByIDAndUserID(context.Context, uint, uint) error
		DeleteByUserID(context.Context, uint) error
		DeleteByTokenHash(context.Context, string) error
	}
)
