package model

import "errors"

var (
	ErrAuthorization = errors.New("you don't have rights for this operation")
	ErrSameCompany   = errors.New("you can't response to ad from your company")

	ErrDatabaseError = errors.New("something wrong with ads database")
	ErrCoreError     = errors.New("something wrong with brm-core service")
)
