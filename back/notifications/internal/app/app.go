package app

import (
	"context"
	"errors"
	"notifications/internal/adapters/grpcstats"
	"notifications/internal/model"
	"notifications/internal/repo"
	"notifications/pkg/logger"
	"time"
)

type appImpl struct {
	r        repo.NotificationsRepo
	statsCli grpcstats.StatsClient
	logs     logger.Logger
}

func (a *appImpl) CreateNotification(ctx context.Context, notification model.Notification) error {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"CompanyId": notification.CompanyId,
			"Method":    "CreateNotification",
		}, err)
	}()
	notification.Date = time.Now().UTC()
	notification.Viewed = false

	err = a.r.CreateNotification(ctx, notification)
	return err
}

func (a *appImpl) GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]model.Notification, uint, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"CompanyId": companyId,
			"Method":    "GetNotifications",
		}, err)
	}()

	notifications, amount, err := a.r.GetNotifications(ctx, companyId, limit, offset, onlyNotViewed)
	return notifications, amount, err
}

func (a *appImpl) GetNotification(ctx context.Context, companyId uint64, notificationId uint64) (model.Notification, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"CompanyId": companyId,
			"Method":    "GetNotification",
		}, err)
	}()

	var notification model.Notification

	notification, err = a.r.GetNotification(ctx, notificationId)
	if err != nil {
		return model.Notification{}, err
	}

	if notification.CompanyId != companyId {
		err = model.ErrPermissionDenied
		return model.Notification{}, err
	}

	return notification, nil
}

func (a *appImpl) SubmitClosedLead(ctx context.Context, companyId uint64, notificationId uint64, submit bool) error {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"CompanyId": companyId,
			"Method":    "SubmitClosedLead",
		}, err)
	}()

	var notification model.Notification
	notification, err = a.r.GetNotification(ctx, notificationId)
	if err != nil {
		return err
	} else if notification.CompanyId != companyId {
		err = model.ErrPermissionDenied
		return err
	} else if notification.Type != model.ClosedLead {
		err = model.ErrWrongNotificationType
	} else if notification.Answered {
		err = model.ErrClosedLeadAlreadyAnswered
		return err
	}

	err = a.r.MarkClosedLeadNotificationAnswered(ctx, notificationId)
	if err != nil {
		return err
	}

	err = a.statsCli.SubmitClosedLead(ctx, notification.ProducerCompany, submit)
	return err
}

func (a *appImpl) writeLog(fields logger.Fields, err error) {
	if errors.Is(err, model.ErrDatabaseError) {
		a.logs.Error(fields, err.Error())
	} else if err != nil {
		a.logs.Info(fields, err.Error())
	} else {
		a.logs.Info(fields, "ok")
	}
}
