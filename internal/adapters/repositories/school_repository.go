package repositories

import (
	"attendance-api/internal/entities"

	"gorm.io/gorm"
)

type (
	SchoolRepository interface {
		Create(school *entities.School) error
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

func (sr *schoolRepository) Create(school *entities.School) error {

	return sr.db.Create(school).Error
}
