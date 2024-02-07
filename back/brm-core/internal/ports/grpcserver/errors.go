package grpcserver

import (
	"brm-core/internal/model"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func mapErrors(err error) error {
	var c codes.Code
	var resErr error

	switch {
	case err == nil:
		return nil
	case errors.Is(err, model.ErrCompanyNotExists):
		c = codes.NotFound
		resErr = model.ErrCompanyNotExists
	case errors.Is(err, model.ErrEmployeeNotExists):
		c = codes.NotFound
		resErr = model.ErrEmployeeNotExists
	case errors.Is(err, model.ErrContactNotExists):
		c = codes.NotFound
		resErr = model.ErrContactNotExists
	case errors.Is(err, model.ErrAuthorization):
		c = codes.PermissionDenied
		resErr = model.ErrAuthorization
	case errors.Is(err, model.ErrDatabaseError):
		c = codes.ResourceExhausted
		resErr = model.ErrDatabaseError
	default:
		c = codes.Unknown
		resErr = model.ErrServiceError
	}

	return status.Errorf(c, resErr.Error())
}
