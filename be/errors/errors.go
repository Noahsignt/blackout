package errors

import (
	"errors"
)

var ErrTooFewPlayers = errors.New("not enough players to start")