package model

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input body or query params")

	ErrCompanyNotExists  = errors.New("company with required id does not exist")
	ErrEmployeeNotExists = errors.New("employee with required id does not exist")
	ErrContactNotExists  = errors.New("contact with required id does not exist")
	ErrIndustryNotExists = errors.New("industry with required id does not exists")
	ErrEmailRegistered   = errors.New("employee with this email is already registered")
	ErrContactExist      = errors.New("this contact already exists")

	ErrAdNotExists = errors.New("ad with required id does not exist")
	ErrSameCompany = errors.New("you can't response to ad from your company")

	ErrPermissionDenied = errors.New("no rights to make this operation")
	ErrUnauthorized     = errors.New("this operation requires authorization")

	ErrAuthError   = errors.New("something wrong with the auth service")
	ErrCoreError   = errors.New("something wrong with the brm-core server")
	ErrCoreUnknown = errors.New("unknown error from brm-core")
	ErrAdsError    = errors.New("something wrong with ads service")
	ErrAdsUnknown  = errors.New("unknown error from brm-ads")
)
