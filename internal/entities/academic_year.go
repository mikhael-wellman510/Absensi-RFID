package entities

import "time"

/*
Updated Entity
*/
type (
	AcademicYear struct {
		Base
		SchoolId  string    `json:"schoolId" gorm:"column:school_id;not null"`
		YearName  string    `json:"yearName" gorm:"column:year_name;not null"`
		StartDate time.Time `json:"startDate" gorm:"column:start_date;not null"`
		EndDate   time.Time `json:"endDate" gorm:"column:end_date;not null"`
		IsActive  bool      `json:"isActive" gorm:"column:is_active;not null"`
		School    School    `gorm:"foreignKey:SchoolId"`
	}

	AcademicYearRequest struct {
		Id        string `json:"id"`
		SchoolId  string `json:"schoolId" binding:"required"`
		YearName  string `json:"yearName" binding:"required"`
		StartDate string `json:"startDate" binding:"required"`
		EndDate   string `json:"endDate" binding:"required"`
	}

	AcademicYearResponse struct {
		Id        string    `json:"id"`
		YearName  string    `json:"yearName"`
		StartDate time.Time `json:"startDate"`
		EndDate   time.Time `json:"endDate"`
		IsActive  bool      `json:"isActive"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)
