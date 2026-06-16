package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	AcademicYearRepository interface {
		Create(ctx context.Context, academicYear *entities.AcademicYear) error
		FindById(ctx context.Context, id string) (*entities.AcademicYear, error)
	}

	academicYearRepository struct {
		db *gorm.DB
	}
)

func NewAcademicYearRepository(db *gorm.DB) AcademicYearRepository {
	return &academicYearRepository{
		db: db,
	}
}

func (a *academicYearRepository) Create(ctx context.Context, academicYear *entities.AcademicYear) error {

	return a.db.WithContext(ctx).Create(academicYear).Error
}

func (a *academicYearRepository) FindById(ctx context.Context, id string) (*entities.AcademicYear, error) {

	academicYear := &entities.AcademicYear{}

	err := a.db.WithContext(ctx).Where("id = ?", id).First(academicYear).Error

	return academicYear, err
}
