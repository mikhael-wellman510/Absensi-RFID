package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	StudentParentRepository interface {
		Create(ctx context.Context, studentParent *entities.StudentParent) error
		FindById(ctx context.Context, id string) (*entities.StudentParent, error)
	}

	studentParentRepository struct {
		db *gorm.DB
	}
)

func NewStudentParentRepository(db *gorm.DB) StudentParentRepository {
	return &studentParentRepository{
		db: db,
	}
}

func (s *studentParentRepository) Create(ctx context.Context, studentParent *entities.StudentParent) error {

	return s.db.WithContext(ctx).Create(studentParent).Error
}

func (s *studentParentRepository) FindById(ctx context.Context, id string) (*entities.StudentParent, error) {

	studentParent := &entities.StudentParent{}

	err := s.db.WithContext(ctx).Where("id = ?", id).First(studentParent).Error

	return studentParent, err
}
