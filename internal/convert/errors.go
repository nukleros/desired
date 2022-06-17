package convert

import (
	"errors"
)

var (
	ErrNilValue = errors.New("error found nil value")

	ErrConvertString  = errors.New("error converting to string")
	ErrConvertInteger = errors.New("error converting to integer")
	ErrConvertFloat   = errors.New("error converting to float64")
	ErrConvertBoolean = errors.New("error converting to boolean")

	ErrConvertSliceString  = errors.New("error convert to slice of strings")
	ErrConvertSliceInteger = errors.New("error convert to slice of integers")
	ErrConvertSliceFloat   = errors.New("error convert to slice of float64s")
	ErrConvertSliceBoolean = errors.New("error convert to slice of booleans")

	ErrConvertMapInterfaceInterface      = errors.New("error converting to map interface to interface type")
	ErrConvertSliceMapInterfaceInterface = errors.New("error converting to slice of map interface to interface type")
)
