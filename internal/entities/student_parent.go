package entities

import (
	"attendance-api/internal/enums"
	"time"
)

/*
Updated Entity
*/
type (
	StudentParent struct {
		Base
		StudentId    string             `json:"studentId" gorm:"column:student_id;not null"`
		ParentId     string             `json:"parentId" gorm:"column:parent_id;not null"`
		RelationType enums.RelationType `json:"relationType" gorm:"column:relation_type;not null"`
		Student      Student            `gorm:"foreignKey:StudentId"`
		Parent       Parent             `gorm:"foreignKey:ParentId"`
	}

	StudentParentRequest struct {
		Id           string             `json:"id"`
		StudentId    string             `json:"studentId" binding:"required"`
		ParentId     string             `json:"parentId" binding:"required"`
		RelationType enums.RelationType `json:"relationType" binding:"required"`
	}

	StudentParentResponse struct {
		Id              string             `json:"id"`
		StudentResponse StudentResponse    `json:"studentResponse"`
		ParentResponse  ParentResponse     `json:"parentResponse"`
		RelationType    enums.RelationType `json:"relationType"`
		CreatedAt       time.Time          `json:"createdAt"`
		UpdatedAt       time.Time          `json:"updatedAt"`
	}
)
