package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	AttendanceEventRepository interface {
		Create(ctx context.Context, attendanceEvent *entities.AttendanceEvent) error
		FindById(ctx context.Context, id string) (*entities.AttendanceEvent, error)
	}

	attendanceEventRepository struct {
		db *gorm.DB
	}
)

func NewAttendanceEventRepository(db *gorm.DB) AttendanceEventRepository {
	return &attendanceEventRepository{
		db: db,
	}
}

func (a *attendanceEventRepository) Create(ctx context.Context, attendanceEvent *entities.AttendanceEvent) error {

	return a.db.WithContext(ctx).Create(attendanceEvent).Error
}

func (a *attendanceEventRepository) FindById(ctx context.Context, id string) (*entities.AttendanceEvent, error) {

	attendanceEvent := &entities.AttendanceEvent{}

	err := a.db.WithContext(ctx).Where("id = ?", id).First(attendanceEvent).Error

	return attendanceEvent, err
}
