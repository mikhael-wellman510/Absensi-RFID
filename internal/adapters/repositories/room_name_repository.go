package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	RoomNameRepository interface {
		Create(ctx context.Context, roomName *entities.RoomName) error
		FindById(ctx context.Context, id string) (*entities.RoomName, error)
	}

	roomNameRepository struct {
		db *gorm.DB
	}
)

func NewRoomNameRepository(db *gorm.DB) RoomNameRepository {

	return &roomNameRepository{
		db: db,
	}
}

func (r *roomNameRepository) Create(ctx context.Context, roomName *entities.RoomName) error {

	return r.db.WithContext(ctx).Create(roomName).Error
}

func (r *roomNameRepository) FindById(ctx context.Context, id string) (*entities.RoomName, error) {

	roomName := &entities.RoomName{}

	err := r.db.WithContext(ctx).Where("id = ?", id).First(roomName).Error

	return roomName, err
}
