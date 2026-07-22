package entities

type AuditLog struct {
	Base
	UserID    *string `json:"userId" gorm:"type:uuid;index"`
	Action    string  `json:"action" gorm:"type:varchar(100);not null"` // e.g., "LOGIN_SUCCESS", "LOGIN_FAILED", "PASSWORD_RESET"
	IpAddress string  `json:"ipAddress" gorm:"type:varchar(45)"`
	UserAgent string  `json:"userAgent" gorm:"type:text"`
	Details   string  `json:"details" gorm:"type:text"`
}
