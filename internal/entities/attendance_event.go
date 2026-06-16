package entities

import (
	"attendance-api/internal/enums"
	"time"
)

/*
Updated Entity
*/
type (
	AttendanceEvent struct {
		Base
		StudentId string                `json:"studentId" gorm:"column:student_id;not null"`
		DeviceId  string                `json:"deviceId" gorm:"column:device_id;not null"`
		EventType enums.EventAttendance `json:"eventType" gorm:"column:event_type;not null"`
		EventTime time.Time             `json:"eventTime" gorm:"column:event_time;not null"`
	}

	AttendanceEventsRequest struct {
		Id        string                `json:"id" binding:"required"`
		StudentId string                `json:"studentId" binding:"required"`
		DeviceId  string                `json:"deviceId" binding:"required"`
		EventType enums.EventAttendance `json:"eventType" binding:"required"`
		EventTime time.Time             `json:"eventTime" binding:"required"`
	}

	AttendanceEventsResponse struct {
		Id              string                `json:"id"`
		StudentResponse StudentResponse       `json:"studentResponse"`
		DeviceResponse  DeviceResponse        `json:"deviceResponse"`
		EventType       enums.EventAttendance `json:"eventType"`
		EventTime       time.Time             `json:"eventTime"`
		CreatedAt       time.Time             `json:"createdAt"`
		UpdatedAt       time.Time             `json:"updatedAt"`
	}
)
