package app

import (
	"context"
	"notifications/internal/model"
)

type App interface {
	CreateNewLeadNotification(ctx context.Context, notification model.Notification) error
	CreateClosedLeadNotification(ctx context.Context, notification model.Notification) error

	GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]model.Notification, error)
	GetNotification(ctx context.Context, companyId uint64, notificationId string) (model.Notification, error)
	SubmitClosedLead(ctx context.Context, companyId uint64, notificationId string, submit bool) error
}
