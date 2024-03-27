package app

import (
	"context"
	"errors"
	"notifications/internal/adapters/grpcstats"
	"notifications/internal/model"
	"notifications/internal/repo"
	"notifications/pkg/logger"
	"strings"
)

type appImpl struct {
	r        repo.Repo
	statsCli grpcstats.StatsClient
	logs     logger.Logger
}

const (
	newLeadNotificationIdPrefix    = "00"
	closedLeadNotificationIdPrefix = "01"
)

func (a *appImpl) CreateNewLeadNotification(ctx context.Context, notification model.Notification) error {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"CompanyId": notification.CompanyId,
			"Method":    "CreateNewLeadNotification",
		}, err)
	}()

	err = a.r.CreateNewLeadNotification(ctx, notification)
	return err
}

func (a *appImpl) CreateClosedLeadNotification(ctx context.Context, notification model.Notification) error {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"CompanyId": notification.CompanyId,
			"Method":    "CreateClosedLeadNotification",
		}, err)
	}()

	err = a.r.CreateClosedLeadNotification(ctx, notification)
	return err
}

func (a *appImpl) GetNotifications(ctx context.Context, companyId uint64, limit uint, offset uint, onlyNotViewed bool) ([]model.Notification, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"CompanyId": companyId,
			"Method":    "GetNotifications",
		}, err)
	}()

	var notifications []model.Notification
	notifications, err = a.r.GetNotifications(ctx, companyId, limit, offset, onlyNotViewed)
	return notifications, err
}

func (a *appImpl) GetNotification(ctx context.Context, companyId uint64, notificationId string) (model.Notification, error) {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"CompanyId": companyId,
			"Method":    "GetNotification",
		}, err)
	}()
	if len(strings.Split(notificationId, "-")) != 2 {
		err = model.ErrInvalidInput
		return model.Notification{}, err
	}

	var notification model.Notification

	switch strings.Split(notificationId, "-")[0] {
	case newLeadNotificationIdPrefix:
		notification, err = a.r.GetNewLeadNotification(ctx, strings.Split(notificationId, "-")[1])
	case closedLeadNotificationIdPrefix:
		notification, err = a.r.GetClosedLeadNotification(ctx, strings.Split(notificationId, "-")[1])
	}
	if err != nil {
		return model.Notification{}, err
	}

	if notification.CompanyId != companyId {
		err = model.ErrPermissionDenied
		return model.Notification{}, err
	}

	return notification, nil
}

func (a *appImpl) SubmitClosedLead(ctx context.Context, companyId uint64, notificationId string, submit bool) error {
	var err error
	defer func() {
		a.writeLog(logger.Fields{
			"CompanyId": companyId,
			"Method":    "SubmitClosedLead",
		}, err)
	}()
	if len(strings.Split(notificationId, "-")) != 2 || strings.Split(notificationId, "-")[0] != closedLeadNotificationIdPrefix {
		err = model.ErrInvalidInput
		return err
	}

	var notification model.Notification
	notification, err = a.r.GetClosedLeadNotification(ctx, strings.Split(notificationId, "-")[1])
	if err != nil {
		return err
	} else if notification.CompanyId != companyId {
		err = model.ErrPermissionDenied
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
