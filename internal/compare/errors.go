package compare

import "errors"

var (
	ErrMismatchedTypes = errors.New("types are mismatched")
	ErrInvalidListType = errors.New("error invalid slice or array type")

	ErrConvertDesired = errors.New("error converting desired")
	ErrConvertActual  = errors.New("error converting actual")
)
