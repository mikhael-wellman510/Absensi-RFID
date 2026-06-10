package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	EducationLevelsRepository interface {
		Create(ctx context.Context, educationLevels *entities.EducationLevel) error
		FindById(ctx context.Context, id string) (*entities.EducationLevel, error)
	}

	educationLevelsRepository struct {
		db *gorm.DB
	}
)

func NewEducationLevelsRepository(db *gorm.DB) EducationLevelsRepository {

	return &educationLevelsRepository{
		db: db,
	}
}

func (e *educationLevelsRepository) Create(ctx context.Context, educationLevels *entities.EducationLevel) error {

	return e.db.WithContext(ctx).Create(educationLevels).Error
}

func (e *educationLevelsRepository) FindById(ctx context.Context, id string) (*entities.EducationLevel, error) {

	educationLevel := &entities.EducationLevel{}

	err := e.db.WithContext(ctx).Where("id = ?", id).First(educationLevel).Error

	return educationLevel, err
}
