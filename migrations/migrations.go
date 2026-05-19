package migrations

import (
	"attendance-api/internal/entities"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entities.School{})

	if err != nil {
		log.Println("Migration Failed")
		return err
	}

	log.Println("Migration Succes")

	return nil
}
