package repo

import (
	"context"
	"notifications/internal/model"
)

type Repo interface {
	CreateNewLeadNotification(ctx context.Context, notification model.Notification) error
	CreateClosedLeadNotification(ctx context.Context, notification model.Notification) error
	GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]model.Notification, error)
	GetNewLeadNotification(ctx context.Context, notificationId string) (model.Notification, error)
	GetClosedLeadNotification(ctx context.Context, notificationId string) (model.Notification, error)
}
