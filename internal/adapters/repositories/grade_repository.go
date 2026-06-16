package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	GradeRepository interface {
		Create(ctx context.Context, grade *entities.Grade) error
		FindById(ctx context.Context, id string) (*entities.Grade, error)
	}

	gradeRepository struct {
		db *gorm.DB
	}
)

func NewGradeRepository(db *gorm.DB) GradeRepository {
	return &gradeRepository{
		db: db,
	}
}

func (g *gradeRepository) Create(ctx context.Context, grade *entities.Grade) error {

	return g.db.WithContext(ctx).Create(grade).Error
}

func (g *gradeRepository) FindById(ctx context.Context, id string) (*entities.Grade, error) {

	grade := &entities.Grade{}

	err := g.db.WithContext(ctx).Where("id = ?", id).First(grade).Error

	return grade, err
}
