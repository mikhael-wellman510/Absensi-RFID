package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	AttendanceRepository interface {
		Create(ctx context.Context, attendance *entities.Attendance) error
		FindById(ctx context.Context, id string) (*entities.Attendance, error)
	}

	attendanceRepository struct {
		db *gorm.DB
	}
)

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{
		db: db,
	}
}

func (a *attendanceRepository) Create(ctx context.Context, attendance *entities.Attendance) error {

	return a.db.WithContext(ctx).Create(attendance).Error
}

func (a *attendanceRepository) FindById(ctx context.Context, id string) (*entities.Attendance, error) {

	attendance := &entities.Attendance{}

	err := a.db.WithContext(ctx).Where("id = ?", id).First(attendance).Error

	return attendance, err
}
