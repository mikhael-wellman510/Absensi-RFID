package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	GradeLevelRepository interface {
		Create(ctx context.Context, gradeLevel *entities.GradeLevel) error
		FindById(ctx context.Context, id string) (*entities.GradeLevel, error)
	}

	gradeLevelRepository struct {
		db *gorm.DB
	}
)

func NewGradeLevelRepository(db *gorm.DB) GradeLevelRepository {
	return &gradeLevelRepository{
		db: db,
	}
}

func (g *gradeLevelRepository) Create(ctx context.Context, gradeLevel *entities.GradeLevel) error {

	return g.db.WithContext(ctx).Create(gradeLevel).Error
}

func (g *gradeLevelRepository) FindById(ctx context.Context, id string) (*entities.GradeLevel, error) {

	gradeLevel := &entities.GradeLevel{}

	err := g.db.WithContext(ctx).Where("id = ?", id).First(gradeLevel).Error

	return gradeLevel, err
}
