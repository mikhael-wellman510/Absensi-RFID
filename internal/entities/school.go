package entities

import "time"

/* Master data */
/*Updated Entity*/
type (
	School struct {
		Base
		Npsn        string `json:"npsn" gorm:"column:npsn;not null"`
		SchoolName  string `json:"schoolName" gorm:"column:school_name;not null;unique"`
		Address     string `json:"address" gorm:"column:address;type:text;not null"`
		SchoolLevel string `json:"schoolLevel" gorm:"column:school_level;not null"`
		Email       string `json:"email" gorm:"column:email;not null"`
		City        string `json:"city" gorm:"column:city;not null"`
		Province    string `json:"province" gorm:"column:province;not null"`
		PhoneNumber string `json:"phoneNumber" gorm:"column:phone_number;not null"`
		IsActive    bool   `json:"isActive" gorm:"column:is_active;not null"`
	}

	SchoolRequest struct {
		Id          string `json:"id"`
		Npsn        string `json:"npsn" binding:"required"`
		SchoolName  string `json:"schoolName" binding:"required"`
		Address     string `json:"address" binding:"required"`
		SchoolLevel string `json:"schoolLevel" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		City        string `json:"city" binding:"required"`
		Province    string `json:"province" binding:"required"`
		PhoneNumber string `json:"phoneNumber" binding:"required"`
	}

	SchoolResponse struct {
		Id          string    `json:"id"`
		Npsn        string    `json:"npsn"`
		SchoolName  string    `json:"schoolName"`
		Address     string    `json:"address"`
		SchoolLevel string    `json:"schoolLevel"`
		Email       string    `json:"email"`
		City        string    `json:"city"`
		Province    string    `json:"province"`
		PhoneNumber string    `json:"phoneNumber"`
		IsActive    bool      `json:"isActive"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}
)
