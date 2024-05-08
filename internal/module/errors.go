package module

import (
	"errors"
)

var (
	ErrInvalidFunction  = errors.New("invalid function definition")
	ErrFunctionNoReturn = errors.New("function has no return values")
	ErrEmptyString      = errors.New("string is empty")
	ErrInvalidArguments = errors.New("invalid function arguments")
	ErrCloneFailed      = errors.New("git clone failed")
)
