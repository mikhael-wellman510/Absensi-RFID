package entities

import "time"

type (
	Teachers struct {
		Base
		Nip         string `json:"nip" gorm:"column:nip;not null"`
		FullName    string `json:"fullName" gorm:"column:full_name;not null"`
		Email       string `json:"email" gorm:"column:email;not null"`
		Address     string `json:"address" gorm:"column:address;type:text;not null"`
		PhoneNumber string `json:"phoneNumber" gorm:"column:phone_number;not null"`
		UserName    string `json:"userName" gorm:"column:user_name;not null"`
		Password    string `json:"password" gorm:"column:password;not null"`
		IsActive    bool   `json:"isActive" gorm:"column:is_active;not null"`
		SchoolId    string `json:"schoolId" gorm:"column:school_id"`
		School      School `gorm:"foreignKey:SchoolId"`
	}

	TeachersRequest struct {
		Id          string `json:"id"`
		Nip         string `json:"nip" binding:"required"`
		FullName    string `json:"fullName" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		Address     string `json:"addres" binding:"required"`
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		UserName    string `json:"userName" binding:"required"`
		Password    string `json:"password" binding:"required"`
		SchoolId    string `json:"schoolId" binding:"required"`
	}

	TeachersResponse struct {
		Id             string         `json:"id"`
		Nip            string         `json:"nip"`
		FullName       string         `json:"fullName"`
		Email          string         `json:"email"`
		Address        string         `json:"address"`
		PhoneNumber    string         `json:"phoneNumber"`
		UserName       string         `json:"userName"`
		Password       string         `json:"password"`
		IsActive       bool           `json:"isActive"`
		SchoolResponse SchoolResponse `json:"schoolResponse"`
		CreatedAt      time.Time      `json:"createdAt"`
		UpdatedAt      time.Time      `json:"updatedAt"`
	}
)
