package model

import "errors"

var (
	ErrInvalidInput   = errors.New("invalid user input")
	ErrWrongFormat    = errors.New("given image has wrong format")
	ErrImageTooBig    = errors.New("given image has very big size")
	ErrImageNotExists = errors.New("image with given id does not exists")

	ErrDatabaseError = errors.New("something wrong with images database")
	ErrServiceError  = errors.New("something wrong with images service")
)
