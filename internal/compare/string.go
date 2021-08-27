package compare

import "errors"

func EqualString(desiredString, actualString interface{}) (bool, error) {
	if desiredString == nil || actualString == nil {
		return false, errors.New(ErrorNilString)
	}

	desiredString, ok := desiredString.(string)
	if !ok {
		return false, errors.New(ErrorConvertString)
	}

	actualString, ok = actualString.(string)
	if !ok {
		return false, errors.New(ErrorConvertString)
	}

	return (desiredString == actualString), nil
}
