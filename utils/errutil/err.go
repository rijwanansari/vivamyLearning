package errutil

import (
	"errors"
)

var (
	ErrRecordNotFound            = errors.New("record not found")
	ErrInvalidInput              = errors.New("invalid input")
	ErrUserIsAlreadyExists       = errors.New("user already exists")
	ErrInvalidLoginCredentials   = errors.New("invalid login credentials")
	ErrAccessTokenSign           = errors.New("error signing access token")
	ErrRefreshTokenSign          = errors.New("error signing refresh token")
	ErrInvalidAuthorizationToken = errors.New("invalid authorization token")
	ErrInvalidJwtSigningMethod   = errors.New("invalid jwt signing method")
	ErrParseJwt                  = errors.New("failed to parse JWT token")
	ErrInvalidAccessToken        = errors.New("invalid access token")
)

func Exists(err error, errs []error) bool {
	for _, e := range errs {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}
