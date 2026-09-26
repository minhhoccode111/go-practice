package persistent

import (
	"context"
	"errors"
	"time"

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

func (r *SessionRepo) Create(ctx context.Context, s *entity.Session) error {
	return r.db.WithContext(ctx).Create(s).Error
}

func (r *SessionRepo) FindByTokenHash(ctx context.Context, hash string) (*entity.Session, error) {
	var s entity.Session
	err := r.db.WithContext(ctx).Where("token_hash = ?", hash).First(&s).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, entity.ErrNotFound
	case err != nil:
		return nil, err
	}
	return &s, nil
}

func (r *SessionRepo) ListByUserID(ctx context.Context, userID uint) ([]entity.Session, error) {
	var sessions []entity.Session
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND expires_at > ?", userID, time.Now()).
		Order("created_at DESC").
		Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *SessionRepo) DeleteByIDAndUserID(ctx context.Context, id, userID uint) error {
	res := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&entity.Session{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return entity.ErrNotFound
	}
	return nil
}

func (r *SessionRepo) DeleteByUserID(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&entity.Session{}).Error
}

func (r *SessionRepo) DeleteByTokenHash(ctx context.Context, hash string) error {
	return r.db.WithContext(ctx).Where("token_hash = ?", hash).Delete(&entity.Session{}).Error
}
