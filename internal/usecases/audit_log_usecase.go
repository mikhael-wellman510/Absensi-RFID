package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"context"
)

type (
	AuditLogService interface {
		logAudit(ctx context.Context, userID *string, action, ip, userAgent, details string)
	}

	auditLogService struct {
		auditLogRepository repositories.AuditLogRepository
	}
)

func NewAuditLogService(auditLogRepository repositories.AuditLogRepository) AuditLogService {
	return &auditLogService{
		auditLogRepository: auditLogRepository,
	}
}
func (al *auditLogService) logAudit(ctx context.Context, userID *string, action, ip, userAgent, details string) {
	_ = al.auditLogRepository.CreateAuditLog(ctx, &entities.AuditLog{
		UserID:    userID,
		Action:    action,
		IpAddress: ip,
		UserAgent: userAgent,
		Details:   details,
	})
}
