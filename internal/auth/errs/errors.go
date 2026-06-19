package errs

import "errors"

var (
	ErrIncorrectPassword = errors.New("wrong password")
	ErrUserNotFound      = errors.New("user not found")
	ErrNotRoot           = errors.New("forbidden. not root")
)
