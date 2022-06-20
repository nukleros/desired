package compare

import (
	"fmt"
	"reflect"
)

func Compare(desiredValue, actualValue interface{}) (bool, error) {
	// if the desired value is nil but the actual value has a value, return a false comparison
	// immediately.  this is the special case to ensuring that types are equal and ensures that
	// if a consumer requests a nil desired object, that the object is in deed nil. alternatively,
	// if the desired value is not nil, but the actual value is nil, return false immediately,
	// otherwise continue processing.
	if isNil(desiredValue) {
		if isNil(actualValue) {
			return true, nil
		}

		return false, nil
	} else {
		if isNil(actualValue) {
			return false, nil
		}
	}

	// ensure that we are attempting to compare equal types, otherwise return an error.
	if !equalTypes(desiredValue, actualValue) {
		return false, fmt.Errorf("%w - desired: [%+T], actual: [%+T]", ErrMismatchedTypes, desiredValue, actualValue)
	}

	switch reflect.TypeOf(desiredValue).Kind() {
	// map types
	case reflect.Map:
		return equalMap(desiredValue, actualValue)
	// list types
	case reflect.Array, reflect.Slice:
		return equalList(desiredValue, actualValue)
	// other types
	default:
		return reflect.DeepEqual(desiredValue, actualValue), nil
	}
}

func equalTypes(desired, actual interface{}) bool {
	return reflect.TypeOf(desired) == reflect.TypeOf(actual)
}

func isNil(value interface{}) bool {
	return fmt.Sprintf("%T", value) == "<nil>"
}
