package model

import "errors"

var (
	ErrInvalidInput = errors.New("given data is invalid")

	ErrPermissionDenied          = errors.New("no rights to make this operation")
	ErrNotificationNotFount      = errors.New("notification with required id does not exist")
	ErrClosedLeadAlreadyAnswered = errors.New("notification of closed lead with required id already answered")

	ErrDatabaseError = errors.New("something wrong with notifications database")
	ErrStatsError    = errors.New("something wrong with stats service")
	ErrServiceError  = errors.New("something wrong with notifications service")
)
