package entities

import (
	"attendance-api/internal/enums"
	"time"
)

/*
Updated Entity
*/
type (
	Grade struct {
		Base
		SchoolLevel enums.SchoolLevel `json:"schoolLevel" gorm:"column:school_level;not null"`
		GradeNumber int               `json:"gradeNumber" gorm:"column:grade_number;not null"`
	}

	GradeRequest struct {
		Id          string            `json:"id"`
		SchoolLevel enums.SchoolLevel `json:"schoolLevel" binding:"required"`
		GradeNumber int               `json:"gradeNumber" binding:"required"`
	}

	GradeResponse struct {
		Id          string            `json:"id"`
		SchoolLevel enums.SchoolLevel `json:"schoolLevel"`
		GradeNumber int               `json:"gradeNumber"`
		CreatedAt   time.Time         `json:"createdAt"`
		UpdatedAt   time.Time         `json:"updatedAt"`
	}
)
