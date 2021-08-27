package compare

import "errors"

func EqualInt(desiredInt, actualInt interface{}) (bool, error) {
	if desiredInt == nil || actualInt == nil {
		return false, errors.New(ErrorNilInt)
	}

	desiredInt, ok := desiredInt.(int)
	if !ok {
		return false, errors.New(ErrorConvertInt)
	}

	actualInt, ok = actualInt.(int)
	if !ok {
		return false, errors.New(ErrorConvertInt)
	}

	return (desiredInt == actualInt), nil
}
