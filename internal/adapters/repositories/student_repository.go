package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	StudentRepository interface {
		Create(ctx context.Context, student *entities.Students) error
		FindById(ctx context.Context, id string) (*entities.Students, error)
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

func (s *studentRepository) Create(ctx context.Context, student *entities.Students) error {

	return s.db.WithContext(ctx).Create(student).Error

}

func (s *studentRepository) FindById(ctx context.Context, id string) (*entities.Students, error) {

	student := &entities.Students{}

	err := s.db.WithContext(ctx).Where("id = ?", id).First(student).Error

	return student, err
}
