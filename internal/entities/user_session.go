package entities

import "time"

type UserSession struct {
	Base
	UserID           string    `json:"userId" gorm:"type:uuid;index;not null"`
	User             User      `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	RefreshTokenHash string    `json:"-" gorm:"column:refresh_token_hash;uniqueIndex;not null"`
	IpAddress        string    `json:"ipAddress" gorm:"column:ip_address"`
	UserAgent        string    `json:"userAgent" gorm:"column:user_agent"`
	IsRevoked        bool      `json:"isRevoked" gorm:"column:is_revoked;default:false"`
	ExpiresAt        time.Time `json:"expiresAt" gorm:"column:expires_at;not null"`
}
