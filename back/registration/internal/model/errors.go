package model

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input body or query params")

	ErrCoreError   = errors.New("something wrong with the brm-core server")
	ErrCoreUnknown = errors.New("unknown error from brm-core")
)
