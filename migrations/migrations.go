package migrations

import (
	"attendance-api/internal/entities"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entities.AcademicYear{},
		&entities.Attendance{},
		&entities.AttendanceEvent{},
		&entities.AuditLog{},
		&entities.ClassRoom{},
		&entities.Device{},
		&entities.Grade{},
		&entities.Parent{},
		&entities.School{},
		&entities.Student{},
		&entities.StudentClass{},
		&entities.StudentParent{},
		&entities.Teacher{},
		&entities.User{},
		&entities.UserSession{},
	)

	if err != nil {
		log.Println("Migration Failed")
		return err
	}

	log.Println("Migration Succes")

	return nil
}
