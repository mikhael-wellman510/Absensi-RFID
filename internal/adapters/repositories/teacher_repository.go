package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	TeacherRepository interface {
		Create(ctx context.Context, teacher *entities.Teacher) error
		FindById(ctx context.Context, id string) (*entities.Teacher, error)
	}

	teacherRepository struct {
		db *gorm.DB
	}
)

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{
		db: db,
	}
}

func (t *teacherRepository) Create(ctx context.Context, teacher *entities.Teacher) error {

	return t.db.WithContext(ctx).Create(teacher).Error
}

func (t *teacherRepository) FindById(ctx context.Context, id string) (*entities.Teacher, error) {

	teacher := &entities.Teacher{}

	err := t.db.WithContext(ctx).Where("id = ?", id).First(teacher).Error

	return teacher, err
}
