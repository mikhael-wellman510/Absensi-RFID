package entities

type (
	EducationLevel struct {
		Base
		Code string `json:"code" gorm:"column:code;not null"`
		Name string `json:"name" gorm:"column:name;not null"`
	}
)
