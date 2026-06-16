package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	StudentClassRepository interface {
		Create(ctx context.Context, studentClass *entities.StudentClass) error
		FindById(ctx context.Context, id string) (*entities.StudentClass, error)
	}

	studentClassRepository struct {
		db *gorm.DB
	}
)

func NewStudentClassRepository(db *gorm.DB) StudentClassRepository {
	return &studentClassRepository{
		db: db,
	}
}

func (s *studentClassRepository) Create(ctx context.Context, studentClass *entities.StudentClass) error {

	return s.db.WithContext(ctx).Create(studentClass).Error
}

func (s *studentClassRepository) FindById(ctx context.Context, id string) (*entities.StudentClass, error) {

	studentClass := &entities.StudentClass{}

	err := s.db.WithContext(ctx).Where("id = ?", id).First(studentClass).Error

	return studentClass, err
}
