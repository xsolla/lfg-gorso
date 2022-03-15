package rsoerror

import (
	"errors"
	"fmt"
)

type Code error

var (
	ErrSystem Code = errors.New("system error")
)

func New(code Code, err error) error {
	return fmt.Errorf(`%e — %e`, ErrSystem, err)
}
