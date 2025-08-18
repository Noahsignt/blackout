package errors

import (
	"errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrTooFewPlayers = errors.New("not enough players to start")
var ErrTooManyPlayers = errors.New("too many players to start")
var ErrDuplicateUsernameOnSignup = errors.New("username already taken")
var ErrPasswordNotLongEnough = errors.New("password not long enough")
var ErrPasswordTooLong = errors.New("password is too long")
var ErrPasswordsDontMatch = errors.New("passwords don't match")