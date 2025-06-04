package errors

import (
	"errors"
)

var ErrTooFewPlayers = errors.New("not enough players to start")
var ErrTooManyPlayers = errors.New("too many players to start")