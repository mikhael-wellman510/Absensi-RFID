package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	TeachersRepository interface {
		Create(ctx context.Context, teachers *entities.Teachers) error
	}

	teachersRepository struct {
		db *gorm.DB
	}
)

func NewTeachersRepository(db *gorm.DB) TeachersRepository {

	return &teachersRepository{db: db}
}
func (t *teachersRepository) Create(ctx context.Context, teachers *entities.Teachers) error {

	return t.db.WithContext(ctx).Create(teachers).Error
}
