package app

import (
	"context"
	"notifications/internal/model"
)

type App interface {
	CreateNotification(ctx context.Context, notification model.Notification) error

	GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]model.Notification, error)
	GetNotification(ctx context.Context, companyId uint64, notificationId uint64) (model.Notification, error)
	SubmitClosedLead(ctx context.Context, companyId uint64, notificationId uint64, submit bool) error
}
