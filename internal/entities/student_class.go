package entities

import "time"

/*
Updated Entity
*/
type (
	StudentClass struct {
		Base
		StudentId      string `json:"studentId" gorm:"column:student_id;not null"`
		ClassRoomId    string `json:"classRoomId" gorm:"column:class_room_id;not null"`
		AcademicYearId string `json:"academicYearId" gorm:"column:academic_year_id;not null"`
	}

	StudentClassRequest struct {
		Id              string `json:"id"`
		StudentId       string `json:"studentId" binding:"required"`
		ClassRoomId     string `json:"classRoomId" binding:"required"`
		AcademicYearsId string `json:"academicYearsId" binding:"required"`
	}

	StudentClassResponse struct {
		Id                   string               `json:"id"`
		StudentResponse      StudentResponse      `json:"studentResponse"`
		ClassRoom            ClassRoomResponse    `json:"classRoomResponse"`
		AcademicYearResponse AcademicYearResponse `json:"academicYearResponse"`
		CreatedAt            time.Time            `json:"createdAt"`
		UpdatedAt            time.Time            `json:"updatedAt"`
	}
)
