package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	FacultiesRepository interface {
		Create(ctx context.Context, faculties *entities.Faculty) error
		FindById(ctx context.Context, id string) (*entities.Faculty, error)
	}

	facultiesRepository struct {
		db *gorm.DB
	}
)

func NewFacultiesRepository(db *gorm.DB) FacultiesRepository {

	return &facultiesRepository{
		db: db,
	}
}

func (r *facultiesRepository) Create(ctx context.Context, faculties *entities.Faculty) error {

	return r.db.WithContext(ctx).Create(faculties).Error
}

func (r *facultiesRepository) FindById(ctx context.Context, id string) (*entities.Faculty, error) {
	faculty := &entities.Faculty{}
	err := r.db.WithContext(ctx).Where("id = ? ", id).First(faculty).Error

	return faculty, err
}
