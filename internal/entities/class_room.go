package entities

import "time"

/*
Updated Entity
*/
type (
	ClassRoom struct {
		Base
		SchoolId        string       `json:"schoolId" gorm:"column:school_id;not null"`
		AcademicYearsId string       `json:"academicYearsId" gorm:"column:academic_years_id;not null"`
		GradeId         string       `json:"gradeId" gorm:"column:grade_id;not null"`
		ClassName       string       `json:"className" gorm:"column:class_name;not null"`
		TeacherId       string       `json:"teacherId" gorm:"column:teacher_id;not null"`
		School          School       `gorm:"foreignKey:SchoolId"`
		AcademicYears   AcademicYear `gorm:"foreignKey:AcademicYearsId"`
		Grade           Grade        `gorm:"foreignKey:GradeId"`
		Teacher         Teacher      `gorm:"foreignKey:TeacherId"`
	}

	ClassRoomRequest struct {
		Id              string `json:"id"`
		SchoolId        string `json:"schoolId" binding:"required"`
		AcademicYearsId string `json:"academicYearsId" binding:"required"`
		GradeId         string `json:"gradeId" binding:"required"`
		ClassName       string `json:"className" binding:"required"`
		TeacherId       string `json:"teacherId" binding:"required"`
	}

	ClassRoomResponse struct {
		Id                    string               `json:"id"`
		SchoolResponse        SchoolResponse       `json:"schoolResponse"`
		AcademicYearsResponse AcademicYearResponse `json:"academicYearsResponse"`
		GradeResponse         GradeResponse        `json:"gradeResponse"`
		TeacherResponse       TeacherResponse      `json:"teacherResponse"`
		ClassName             string               `json:"className"`
		CreatedAt             time.Time            `json:"createdAt"`
		UpdatedAt             time.Time            `json:"updatedAt"`
	}
)
