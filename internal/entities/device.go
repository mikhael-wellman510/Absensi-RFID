package entities

import "time"

/*
Updated Entity
*/
type (
	Device struct {
		Base
		SchoolId  string `json:"schoolId" gorm:"column:school_id;not null"`
		Code      string `json:"code" gorm:"column:code;not null"`
		Name      string `json:"name" gorm:"column:name;not null"`
		Location  string `json:"location" gorm:"column:location;not null"`
		IpAddress string `json:"ipAddress" gorm:"column:ip_address;not null"`
		IsActive  bool   `json:"isActive" gorm:"column:is_active;not null"`
		School    School `gorm:"foreignKey:SchoolId"`
	}

	DeviceRequest struct {
		Id        string `json:"id" binding:"required"`
		SchoolId  string `json:"schoolId" binding:"required"`
		Code      string `json:"code" binding:"required"`
		Name      string `json:"name" binding:"required"`
		Location  string `json:"location" binding:"required"`
		IpAddress string `json:"ipAddress" binding:"required"`
	}

	DeviceResponse struct {
		Id             string         `json:"id"`
		SchoolResponse SchoolResponse `json:"schoolResponse"`
		Code           string         `json:"code"`
		Name           string         `json:"name"`
		Location       string         `json:"location"`
		IpAddress      string         `json:"ipAddress"`
		CreatedAt      time.Time      `json:"createdAt"`
		UpdatedAt      time.Time      `json:"updatedAt"`
	}
)
