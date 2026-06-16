package entities

import (
	"attendance-api/internal/enums"
	"time"
)

/*
Entity Updated
*/
type (
	Attendance struct {
		Base
		StudentId        string                 `json:"studentId" gorm:"column:student_id;not null"`
		AttendanceDate   time.Time              `json:"attendanceDate" gorm:"column:attendance_date;not null"`
		FirstCheckIn     time.Time              `json:"firstCheckIn" gorm:"column:first_check_in;not null"`
		LastCheckIn      time.Time              `json:"lastCheckIn" gorm:"column:last_check_in;not null"`
		StatusAttendance enums.StatusAttendance `json:"statusAttendance" gorm:"column:status_attendance;not null"`
		Student          Student                `gorm:"foreignKey:StudentId"`
	}

	AttendanceRequest struct {
		Id               string                 `json:"id"`
		StudentId        string                 `json:"studentId" binding:"required"`
		AttendanceDate   time.Time              `json:"attendanceDate" binding:"required"`
		FirstCheckIn     time.Time              `json:"firstCheckIn" binding:"required"`
		LastCheckIn      time.Time              `json:"lastCheckIn" binding:"required"`
		StatusAttendance enums.StatusAttendance `json:"statusAttendance" binding:"required"`
	}

	AttendanceResponse struct {
		Id               string                 `json:"id"`
		StudentResponse  StudentResponse        `json:"studentResponse"`
		AttendanceDate   time.Time              `json:"attendanceDate"`
		FirstCheckIn     time.Time              `json:"firstCheckIn"`
		LastCheckIn      time.Time              `json:"lastCheckIn"`
		StatusAttendance enums.StatusAttendance `json:"statusAttendance"`
		CreatedAt        time.Time              `json:"createdAt"`
		UpdatedAt        time.Time              `json:"updatedAt"`
	}
)
