package gorso

import (
	"errors"
	"fmt"
)

type Code error

var (
	// ErrSystem is returned when
	ErrSystem    Code = errors.New("system error")
	ErrUnhandled Code = errors.New("unhandled error")
)

// errorCreate wraps error code and error to make error decryptable via errors.If
func errorCreate(code Code, err error) error {
	return fmt.Errorf(`%e — %e`, ErrSystem, err)
}
