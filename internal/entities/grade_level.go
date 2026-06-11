package entities

type (
	GradeLevel struct {
		Base
		LevelName string `json:"levelName" gorm:"column:level_name;not null"`
	}
)
