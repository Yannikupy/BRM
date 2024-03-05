package model

import "errors"

var (
	ErrLeadNotExists     = errors.New("lead with required id does not exist")
	ErrCompanyNotExists  = errors.New("company with required id does not exist")
	ErrEmployeeNotExists = errors.New("employee with required id does not exist")
	ErrAdNotExists       = errors.New("ad with required id does not exist")
	ErrAuthorization     = errors.New("no rights to make operation")

	ErrDatabaseError = errors.New("something wrong with the database")
	ErrServiceError  = errors.New("something wrong with the server")
	ErrCoreError     = errors.New("something wrong with the brm-core server")
	ErrAdsError      = errors.New("something wrong with ads service")
)
