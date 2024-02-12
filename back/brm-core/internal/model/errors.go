package model

import "errors"

var ErrCompanyNotExists = errors.New("company with required id does not exist")
var ErrEmployeeNotExists = errors.New("employee with required id does not exist")
var ErrContactNotExists = errors.New("contact with required id does not exist")

var ErrAuthorization = errors.New("no rights to make operation: ownerId mismatched")

var ErrDatabaseError = errors.New("something wrong with the database")
var ErrServiceError = errors.New("something wrong with the server")
