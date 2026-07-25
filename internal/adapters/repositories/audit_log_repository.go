package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	AuditLogRepository interface {
		CreateAuditLog(context.Context, *entities.AuditLog) error
	}

	auditLogRepository struct {
		db *gorm.DB
	}
)

func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{
		db: db,
	}
}

func (al *auditLogRepository) CreateAuditLog(ctx context.Context, log *entities.AuditLog) error {
	return al.db.WithContext(ctx).Create(log).Error
}
