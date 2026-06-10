package entities

import "time"

type (
	EducationLevel struct {
		Base
		Code string `json:"code" gorm:"column:code;not null"`
		Name string `json:"name" gorm:"column:name;not null"`
	}

	EducationLevelsRequest struct {
		Id   string `json:"id"`
		Code string `json:"code" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	EducationLevelsResponse struct {
		Id        string    `json:"id"`
		Code      string    `json:"code"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)
