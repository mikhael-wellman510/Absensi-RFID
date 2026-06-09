package entities

type (
	StudyProgram struct {
		Base
		FacultyId string  `json:"facultyId" gorm:"column:faculty_id;not null"`
		Code      string  `json:"code" gorm:"column:code;not null"`
		Name      string  `json:"name" gorm:"column:name;not null"`
		Degree    int     `json:"degree" gorm:"column:degree;not null"`
		Faculty   Faculty `gorm:"foreigner:FacultyId"`
	}
)
