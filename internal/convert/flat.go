package convert

import (
	"fmt"
)

// ToString converts an interface to a string and returns an error if unable to.
func ToString(in interface{}) (out string, err error) {
	if in == nil {
		return out, ErrNilValue
	}

	out, ok := in.(string)
	if !ok {
		return out, fmt.Errorf("%w for value %v", ErrConvertString, in)
	}

	return out, nil
}

// ToInteger converts an interface to an integer and returns an error if unable to.
func ToInteger(in interface{}) (out int, err error) {
	if in == nil {
		return out, ErrNilValue
	}

	out, ok := in.(int)
	if !ok {
		return out, fmt.Errorf("%w for value %v", ErrConvertInteger, in)
	}

	return out, nil
}

// ToFloat converts an interface to a float64 and returns an error if unable to.
func ToFloat(in interface{}) (out float64, err error) {
	if in == nil {
		return out, ErrNilValue
	}

	out, ok := in.(float64)
	if !ok {
		return out, fmt.Errorf("%w for value %v", ErrConvertFloat, in)
	}

	return out, nil
}

// ToBoolean converts an interface to a boolean and returns an error if unable to.
func ToBoolean(in interface{}) (out bool, err error) {
	if in == nil {
		return out, ErrNilValue
	}

	out, ok := in.(bool)
	if !ok {
		return out, fmt.Errorf("%w for value %v", ErrConvertBoolean, in)
	}

	return out, nil
}

// BooleanToInteger convert a boolean type to an integer.
func BooleanToInteger(in bool) (out int) {
	if in {
		return 1
	}

	return out
}
