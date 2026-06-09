package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	SchoolRepository interface {
		Create(ctx context.Context, school *entities.School) error
		FindById(ctx context.Context, id string) (*entities.School, error)
	}

	schoolRepository struct {
		db *gorm.DB
	}
)

func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &schoolRepository{
		db: db,
	}
}

func (sr *schoolRepository) Create(ctx context.Context, school *entities.School) error {

	return sr.db.WithContext(ctx).Create(school).Error
}

func (sr *schoolRepository) FindById(ctx context.Context, id string) (*entities.School, error) {
	schools := &entities.School{}

	err := sr.db.WithContext(ctx).Where("id = ?", id).First(schools).Error
	return schools, err
}
