package entities

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (s *Base) BeforeCreate(tx *gorm.DB) error {
	log.Println("Call before Create")
	s.ID = uuid.New().String()

	log.Println("Uuid : ", s.ID)
	return nil
}
