package compare

import (
	"fmt"
	"reflect"
)

func Compare(desiredValue, actualValue interface{}) (bool, error) {

	if passNilComparison := compareNil(desiredValue, actualValue); !passNilComparison {
		return false, nil
	}

	if !equalTypes(desiredValue, actualValue) {
		return false, fmt.Errorf("%w\n\ndesired: %+T\n\nactual: %+T", ErrMismatchedTypes, desiredValue, actualValue)
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

func compareNil(desired, actual interface{}) bool {
	// treat a nil desired value as a value that is not checked for a desirable condition
	if desired == nil {
		return true
	}

	// if the actual value is nil, we are not in a desired state since we have already confirmed
	// above that the desired value is not nil
	return actual != nil
}
