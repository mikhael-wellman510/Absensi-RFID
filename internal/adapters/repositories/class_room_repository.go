package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	ClassRoomRepository interface {
		Create(ctx context.Context, classRoom *entities.ClassRoom) error
		FindById(ctx context.Context, id string) (*entities.ClassRoom, error)
	}

	classRoomRepository struct {
		db *gorm.DB
	}
)

func NewClassRoomRepository(db *gorm.DB) ClassRoomRepository {
	return &classRoomRepository{
		db: db,
	}
}

func (c *classRoomRepository) Create(ctx context.Context, classRoom *entities.ClassRoom) error {

	return c.db.WithContext(ctx).Create(classRoom).Error
}

func (c *classRoomRepository) FindById(ctx context.Context, id string) (*entities.ClassRoom, error) {

	classRoom := &entities.ClassRoom{}

	err := c.db.WithContext(ctx).Where("id = ?", id).First(classRoom).Error

	return classRoom, err
}
