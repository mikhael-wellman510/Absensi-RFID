package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	UserSessionRepository interface {
		CreateSession(ctx context.Context, session *entities.UserSession) error
		FindSessionByHash(ctx context.Context, hash string) (*entities.UserSession, error)
		RevokeSessionByID(ctx context.Context, sessionID string) error
		RevokeAllUserSessions(ctx context.Context, userID string) error
	}

	userSessionRepository struct {
		db *gorm.DB
	}
)

func NewUserSessionRepository(db *gorm.DB) UserSessionRepository {
	return &userSessionRepository{
		db: db,
	}
}

func (r *userSessionRepository) CreateSession(ctx context.Context, session *entities.UserSession) error {
	return r.db.WithContext(ctx).Create(session).Error
}

func (r *userSessionRepository) FindSessionByHash(ctx context.Context, hash string) (*entities.UserSession, error) {
	session := &entities.UserSession{}
	err := r.db.WithContext(ctx).Where("refresh_token_hash = ?", hash).First(session).Error
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (r *userSessionRepository) RevokeSessionByID(ctx context.Context, sessionID string) error {
	return r.db.WithContext(ctx).
		Model(&entities.UserSession{}).
		Where("id = ?", sessionID).
		Update("is_revoked", true).Error
}

func (r *userSessionRepository) RevokeAllUserSessions(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).
		Model(&entities.UserSession{}).
		Where("user_id = ?", userID).
		Update("is_revoked", true).Error
}
