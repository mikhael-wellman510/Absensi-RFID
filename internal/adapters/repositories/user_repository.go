package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(ctx context.Context, user *entities.User) error
		FindById(ctx context.Context, id string) (*entities.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *entities.User) error {

	return u.db.WithContext(ctx).Create(user).Error
}

func (u *userRepository) FindById(ctx context.Context, id string) (*entities.User, error) {

	user := &entities.User{}

	err := u.db.WithContext(ctx).Where("id = ?", id).First(user).Error

	return user, err
}
