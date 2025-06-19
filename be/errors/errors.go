package errors

import (
	"errors"
)

var ErrTooFewPlayers = errors.New("not enough players to start")
var ErrTooManyPlayers = errors.New("too many players to start")
var ErrDuplicateUsernameOnSignup = errors.New("username already taken")
var PasswordNotLongEnough = errors.New("password not long enough")
var PasswordTooLong = errors.New("password is too long")