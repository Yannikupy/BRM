package model

import "errors"

var (
	ErrInvalidInput      = errors.New("invalid input body or query params")
	ErrIndustryNotExists = errors.New("industry with required id does not exist")

	ErrCoreError   = errors.New("something wrong with the brm-core server")
	ErrCoreUnknown = errors.New("unknown error from brm-core")
)
