package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"context"
)

type (
	UserSessionService interface {
		CreateSession(ctx context.Context, session *entities.UserSession) error
		FindSessionByHash(ctx context.Context, hash string) (*entities.UserSession, error)
		RevokeSessionByID(ctx context.Context, sessionID string) error
		RevokeAllUserSessions(ctx context.Context, userID string) error
	}

	userSessionService struct {
		userSessionRepository repositories.UserSessionRepository
	}
)

func NewUserSessionService(userSessionRepository repositories.UserSessionRepository) UserSessionService {

	return &userSessionService{
		userSessionRepository: userSessionRepository,
	}
}
func (u *userSessionService) CreateSession(ctx context.Context, session *entities.UserSession) error {

	return u.userSessionRepository.CreateSession(ctx, session)
}

func (u *userSessionService) FindSessionByHash(ctx context.Context, hash string) (*entities.UserSession, error) {
	return u.userSessionRepository.FindSessionByHash(ctx, hash)
}

func (u *userSessionService) RevokeSessionByID(ctx context.Context, sessionID string) error {

	return u.userSessionRepository.RevokeSessionByID(ctx, sessionID)
}

func (u *userSessionService) RevokeAllUserSessions(ctx context.Context, userID string) error {

	return u.userSessionRepository.RevokeAllUserSessions(ctx, userID)
}
