package compare

import (
	"reflect"
)

func Compare(desiredValue, actualValue interface{}) (bool, error) {

	if !equalTypes(desiredValue, actualValue) {
		return false, ErrMismatchedTypes
	}

	switch reflect.ValueOf(desiredValue).Kind() {
	// map types
	case reflect.Map:
		return EqualMap(desiredValue, actualValue)
	// list types
	case reflect.Array, reflect.Slice:
		return EqualList(desiredValue, actualValue)
	// other types
	default:
		return reflect.DeepEqual(desiredValue, actualValue), nil
	}
}

func equalTypes(desired, actual interface{}) bool {
	return reflect.TypeOf(desired) == reflect.TypeOf(actual)
}
