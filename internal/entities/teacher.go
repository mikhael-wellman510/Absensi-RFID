package entities

import (
	"attendance-api/internal/enums"
	"time"
)

/* Updated entity*/
type (
	Teacher struct {
		Base
		UserId    string       `json:"userId" gorm:"column:user_id;not null"`
		SchoolId  string       `json:"schoolId" gorm:"column:school_id;not null"`
		Nip       string       `json:"nip" gorm:"column:nip;not null"`
		Gender    enums.Gender `json:"gender" gorm:"column:gender;not null"`
		BirthDate time.Time    `json:"birthDate" gorm:"column:birth_date;not null"`
		Address   string       `json:"address" gorm:"column:address;type:text;not null"`
		IsActive  bool         `json:"isActive" gorm:"column:is_active;not null"`
		User      User         `gorm:"foreignKey:userId"`
		School    School       `gorm:"foreignKey:SchoolId"`
	}

	TeacherRequest struct {
		Id        string       `json:"id"`
		UserId    string       `json:"userId" binding:"required"`
		SchoolId  string       `json:"schoolId" binding:"required"`
		Nip       string       `json:"nip" binding:"required"`
		Gender    enums.Gender `json:"gender" binding:"required"`
		BirthDate string       `json:"birthDate" binding:"required"`
		Address   string       `json:"address" binding:"required"`
	}

	TeacherResponse struct {
		Id             string         `json:"id"`
		Nip            string         `json:"nip"`
		Gender         enums.Gender   `json:"gender"`
		BirthDate      time.Time      `json:"birthDate"`
		Address        string         `json:"address"`
		IsActive       bool           `json:"isActive"`
		SchoolResponse SchoolResponse `json:"schoolResponse"`
		UserResponse   UserResponse   `json:"userResponse"`
		CreatedAt      time.Time      `json:"createdAt"`
		UpdatedAt      time.Time      `json:"updatedAt"`
	}
)
