package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	StudentRepository interface {
		Create(ctx context.Context, student *entities.Student) error
		FindById(ctx context.Context, id string) (*entities.Student, error)
	}

	studentRepository struct {
		db *gorm.DB
	}
)

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{
		db: db,
	}
}

func (s *studentRepository) Create(ctx context.Context, student *entities.Student) error {

	return s.db.WithContext(ctx).Create(student).Error
}

func (s *studentRepository) FindById(ctx context.Context, id string) (*entities.Student, error) {

	student := &entities.Student{}

	err := s.db.WithContext(ctx).Preload("School").First(student, "id=?", id).Error

	return student, err
}
