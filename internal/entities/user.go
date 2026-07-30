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

		// Field Keamanan Enterprise
		IsEmailVerified        bool       `json:"isEmailVerified" gorm:"column:is_email_verified;default:false"`
		EmailVerificationToken string     `json:"-" gorm:"column:email_verification_token"`
		FailedLoginAttempts    int        `json:"-" gorm:"column:failed_login_attempts;default:0"`
		LockedUntil            *time.Time `json:"-" gorm:"column:locked_until;default:null"`
		PasswordResetToken     string     `json:"-" gorm:"column:password_reset_token"`
		PasswordResetExpires   *time.Time `json:"-" gorm:"column:password_reset_expires;default:null"`
	}

	UserRequest struct {
		Id          string     `json:"id"`
		FullName    string     `json:"fullName" binding:"required"`
		Email       string     `json:"email" binding:"required"`
		PhoneNumber string     `json:"phoneNumber" binding:"required"`
		Password    string     `json:"password" binding:"required"`
		Role        enums.Role `json:"role" binding:"required"`
	}

	// CreateUserRequest digunakan untuk menerima data pendaftaran/pembuatan user baru
	CreateUserRequest struct {
		FullName    string     `json:"fullName" binding:"required"`
		Email       string     `json:"email" binding:"required,email"`
		PhoneNumber string     `json:"phoneNumber" binding:"required"`
		Password    string     `json:"password" binding:"required,min=8"`
		Role        enums.Role `json:"role" binding:"required"`
	}

	UserResponse struct {
		Id              string     `json:"id"`
		FullName        string     `json:"fullName"`
		Email           string     `json:"email"`
		PhoneNumber     string     `json:"phoneNumber"`
		Password        string     `json:"password"`
		Role            enums.Role `json:"role"`
		LastLogin       time.Time  `json:"lastLogin"`
		IsActive        bool       `json:"isActive"`
		IsEmailVerified bool       `json:"isEmailVerified"`
		CreatedAt       time.Time  `json:"createdAt"`
		UpdatedAt       time.Time  `json:"updatedAt"`
	}

	// LoginRequest digunakan untuk menerima payload autentikasi awal user
	LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	UserProfile struct {
		Id          string     `json:"id"`
		FullName    string     `json:"fullName"`
		Email       string     `json:"email"`
		PhoneNumber string     `json:"phoneNumber"`
		Role        enums.Role `json:"role"`
	}

	// AuthResponse digunakan sebagai format response token ke client setelah login/refresh
	AuthResponse struct {
		AccessToken  string      `json:"accessToken"`
		RefreshToken string      `json:"refreshToken"`
		TokenType    string      `json:"tokenType"`
		UserProfile  UserProfile `json:"userProfile"`
	}

	// RefreshTokenRequest digunakan untuk meminta access token baru dengan token rotasi
	RefreshTokenRequest struct {
		RefreshToken string `json:"refreshToken" binding:"required"`
	}

	// ForgotPasswordRequest digunakan untuk menerima permintaan link/token reset password
	ForgotPasswordRequest struct {
		Email string `json:"email" binding:"required,email"`
	}

	// ResetPasswordRequest digunakan untuk memproses pembaharuan password dengan token
	ResetPasswordRequest struct {
		Token       string `json:"token" binding:"required"`
		NewPassword string `json:"newPassword" binding:"required,min=8"`
	}
)
