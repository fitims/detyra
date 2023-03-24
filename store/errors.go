package store

import "errors"

var (
	UserDoesNotExistErr      = errors.New("user does not exist")
	UserAlreadyRegisteredErr = errors.New("user is registered already")
)
