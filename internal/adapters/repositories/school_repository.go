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

func (s *schoolRepository) Create(ctx context.Context, school *entities.School) error {

	return s.db.WithContext(ctx).Create(school).Error
}

func (s *schoolRepository) FindById(ctx context.Context, id string) (*entities.School, error) {

	school := &entities.School{}

	err := s.db.WithContext(ctx).Where("id = ?", id).First(school).Error

	return school, err
}
