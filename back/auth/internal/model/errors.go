package model

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input body or query params")

	ErrWrongPassword      = errors.New("wrong password of the user with required email")
	ErrCreateAccessToken  = errors.New("unable to create access token")
	ErrParsingAccessToken = errors.New("unable to parse access token")

	ErrAccessTokenNotExpired = errors.New("existing access token has not expired yet")
	ErrTokensNotExist        = errors.New("required tokens are not exist")

	ErrEmployeeNotExists = errors.New("employee with required email not exists")

	ErrAuthRepoError = errors.New("something wrong with tokens repository")
	ErrPassRepoError = errors.New("something wrong with passwords repository")

	ErrServiceError = errors.New("something wrong with the server")
)
