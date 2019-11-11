package db

import (
	"errors"
)

type Persist interface {
	test()
}

var ErrUserNotFound = errors.New("db: user not found")
