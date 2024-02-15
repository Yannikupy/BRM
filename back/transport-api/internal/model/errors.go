package model

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input body or query params")

	ErrCompanyNotExists  = errors.New("company with required id does not exist")
	ErrEmployeeNotExists = errors.New("employee with required id does not exist")
	ErrContactNotExists  = errors.New("contact with required id does not exist")

	ErrPermissionDenied = errors.New("no rights to make this operation")
	ErrUnauthorized     = errors.New("this operation requires authorization")

	ErrCoreError   = errors.New("something wrong with the brm-core server")
	ErrCoreUnknown = errors.New("unknown error from brm-core")
)
