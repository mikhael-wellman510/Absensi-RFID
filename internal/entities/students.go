package entities

import "attendance-api/internal/enums"

type (
	Students struct {
		Base
		StudentNumber    string         `json:"studentNumber" gorm:"column:student_number;not null;unique"`
		FullName         string         `json:"fullName" gorm:"column:full_name;not null"`
		Gender           enums.Gender   `json:"gender" gorm:"type:varchar(10);column:gender;not null"`
		Status           bool           `json:"status" gorm:"column:status;not null"`
		EducationLevelId string         `json:"educationLevelId" gorm:"column:education_level_id;not null"`
		StudyProgramId   string         `json:"studyProgramId" gorm:"column:study_program_id;not null"`
		EducationLevel   EducationLevel `gorm:"foreignKey:EducationLevelId"`
		StudyProgram     StudyProgram   `gorm:"foreignKey:StudyProgramId"`
	}
)
