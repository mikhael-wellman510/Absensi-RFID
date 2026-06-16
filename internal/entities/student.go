package entities

import (
	"attendance-api/internal/enums"
	"time"
)

/*
Updated Entity
*/
type (
	Student struct {
		Base
		SchoolId       string       `json:"schoolId" gorm:"column:school_id;not null"`
		Nis            string       `json:"nis" gorm:"column:nis;not null"`
		Nisn           string       `json:"nisn" gorm:"column:nisn;not null"`
		FullName       string       `json:"fullName" gorm:"column:full_name;not null"`
		Gender         enums.Gender `json:"gender" gorm:"column:gender;not null"`
		BirthDate      time.Time    `json:"birthDate" gorm:"column:birth_date;not null"`
		Address        string       `json:"address" gorm:"column:address;not null"`
		EnrollmentDate time.Time    `json:"enrollmentDate" gorm:"column:enrollment_date;not null"`
		Status         enums.Status `json:"status" gorm:"column:status;not null"`
		School         School       `gorm:"foreignKey:SchoolId"`
	}

	StudentRequest struct {
		Id             string       `json:"id"`
		SchoolId       string       `json:"schoolId" binding:"required"`
		Nis            string       `json:"nis" binding:"required"`
		Nisn           string       `json:"nisn" binding:"required"`
		FullName       string       `json:"fullName" binding:"required"`
		Gender         enums.Gender `json:"gender" binding:"required"`
		BirthDate      string       `json:"birthDate" binding:"required"`
		Address        string       `json:"address" binding:"required"`
		EnrollmentDate string       `json:"enrollmentDate" binding:"required"`
		Status         enums.Status `json:"status" binding:"required"`
	}

	StudentResponse struct {
		Id             string         `json:"id"`
		Nis            string         `json:"nis"`
		Nisn           string         `json:"nisn"`
		FullName       string         `json:"fullName"`
		Gender         enums.Gender   `json:"gender"`
		BirthDate      time.Time      `json:"birthDate"`
		Address        string         `json:"address"`
		EnrollmentDate time.Time      `json:"enrollmentDate"`
		Status         enums.Status   `json:"status"`
		SchoolResponse SchoolResponse `json:"schoolResponse"`
		CreatedAt      time.Time      `json:"createdAt"`
		UpdatedAt      time.Time      `json:"updatedAt"`
	}
)
