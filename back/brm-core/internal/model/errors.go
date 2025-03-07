package model

import "errors"

var ErrCompanyNotExists = errors.New("company with required id does not exist")
var ErrEmployeeNotExists = errors.New("employee with required id does not exist")
var ErrContactNotExists = errors.New("contact with required id does not exist")
var ErrIndustryNotExists = errors.New("industry with required id does not exist")

var ErrEmailRegistered = errors.New("employee with this email is already registered")
var ErrContactExist = errors.New("this contact already exist")

var ErrAuthorization = errors.New("no rights to make operation: ownerId mismatched")

var ErrDatabaseError = errors.New("something wrong with the database")
var ErrServiceError = errors.New("something wrong with the server")
var ErrAuthServiceError = errors.New("something wrong with auth service")
