package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	DeviceRepository interface {
		Create(ctx context.Context, device *entities.Device) error
		FindById(ctx context.Context, id string) (*entities.Device, error)
	}

	deviceRepository struct {
		db *gorm.DB
	}
)

func NewDeviceRepository(db *gorm.DB) DeviceRepository {
	return &deviceRepository{
		db: db,
	}
}

func (d *deviceRepository) Create(ctx context.Context, device *entities.Device) error {

	return d.db.WithContext(ctx).Create(device).Error
}

func (d *deviceRepository) FindById(ctx context.Context, id string) (*entities.Device, error) {

	device := &entities.Device{}

	err := d.db.WithContext(ctx).Where("id = ?", id).First(device).Error

	return device, err
}
