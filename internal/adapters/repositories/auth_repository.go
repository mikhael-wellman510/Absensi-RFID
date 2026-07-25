package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	AuthRepository interface {
		FindByEmail(ctx context.Context, email string) (*entities.User, error)
		FindByID(ctx context.Context, id string) (*entities.User, error)
		UpdateUser(ctx context.Context, user *entities.User) error
		Create(ctx context.Context, user *entities.User) error
		FindByPhoneNumber(ctx context.Context, phone string) (*entities.User, error)

		// FindByResetToken Password Reset & Email Verification
		FindByResetToken(ctx context.Context, token string) (*entities.User, error)
		FindByVerificationToken(ctx context.Context, token string) (*entities.User, error)
	}

	authRepository struct {
		db *gorm.DB
	}
)

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.WithContext(ctx).Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) FindByID(ctx context.Context, id string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *authRepository) FindByResetToken(ctx context.Context, token string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.WithContext(ctx).Where("password_reset_token = ?", token).First(user).Error
	return user, err
}

func (r *authRepository) FindByVerificationToken(ctx context.Context, token string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.WithContext(ctx).Where("email_verification_token = ?", token).First(user).Error
	return user, err
}

func (r *authRepository) Create(ctx context.Context, user *entities.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *authRepository) FindByPhoneNumber(ctx context.Context, phone string) (*entities.User, error) {
	user := &entities.User{}
	err := r.db.WithContext(ctx).Where("phone_number = ?", phone).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
