package entities

import (
	"attendance-api/internal/enums"
	"time"
)

/*Updated Entity*/
type (
	User struct {
		Base
		FullName    string     `json:"fullName" gorm:"column:full_name;not null"`
		Email       string     `json:"email" gorm:"column:email;not null"`
		PhoneNumber string     `json:"phoneNumber" gorm:"column:phone_number;not null"`
		Password    string     `json:"password" gorm:"column:password;not null"`
		Role        enums.Role `json:"role" gorm:"column:role;not null"`
		LastLogin   time.Time  `json:"lastLogin" gorm:"column:last_login;default:null"`
		IsActive    bool       `json:"isActive" gorm:"column:is_active"`
	}

	UserRequest struct {
		Id          string     `json:"id"`
		FullName    string     `json:"fullName" binding:"required"`
		Email       string     `json:"email" binding:"required"`
		PhoneNumber string     `json:"phoneNumber" binding:"required"`
		Password    string     `json:"password" binding:"required"`
		Role        enums.Role `json:"role" binding:"required"`
	}

	UserResponse struct {
		Id          string     `json:"id"`
		FullName    string     `json:"fullName"`
		Email       string     `json:"email"`
		PhoneNumber string     `json:"phoneNumber"`
		Password    string     `json:"password"`
		Role        enums.Role `json:"role"`
		LastLogin   time.Time  `json:"lastLogin"`
		IsActive    bool       `json:"isActive"`
		CreatedAt   time.Time  `json:"createdAt"`
		UpdatedAt   time.Time  `json:"updatedAt"`
	}
)
