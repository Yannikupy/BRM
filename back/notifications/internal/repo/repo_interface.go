package repo

import (
	"context"
	"notifications/internal/model"
)

type NotificationsRepo interface {
	CreateNotification(ctx context.Context, notification model.Notification) error
	GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]model.Notification, error)
	GetNotification(ctx context.Context, id uint64) (model.Notification, error)
	MarkClosedLeadNotificationAnswered(ctx context.Context, notificationId uint64) error
}
