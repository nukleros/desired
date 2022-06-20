package compare

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/nukleros/desired/internal/convert"
)

// equalList takes in a list of any type and returns the result of the comparison.
func equalList(desiredList, actualList interface{}) (bool, error) {
	desiredAsValue := reflect.ValueOf(desiredList)
	actualAsValue := reflect.ValueOf(actualList)

	// return equality if length of both lists are not the same
	if desiredAsValue.Len() != actualAsValue.Len() {
		return true, nil
	}

	// handle a map type first since ordering is important and potentially unpredictable with maps
	if desiredAsValue.Index(0).Kind() == reflect.Map {
		return equalSliceMapInterfaceInterface(desiredList, actualList)
	}

	switch desiredAsType := desiredList.(type) {
	case []string:
		return equalSliceStringInterface(desiredList, actualList)
	case []int, []int8, []int16, []int32, []int64:
		return equalSliceIntegerInterface(desiredList, actualList)
	case []float32, []float64:
		return equalSliceFloatInterface(desiredList, actualList)
	case []bool:
		return equalSliceBooleanInterface(desiredList, actualList)
	case []interface{}:
		return equalSliceInterface(desiredAsType, actualList.([]interface{}))
	}

	return false, fmt.Errorf("%w of type %T", ErrInvalidListType, desiredList)
}

// equalSliceStringInterface takes in an expected slice of string type, as an interface, converts
// it appropriately and returns the result of the comparison.
func equalSliceStringInterface(desiredSlice, actualSlice interface{}) (bool, error) {
	desired, err := convert.ToSliceString(desiredSlice)
	if err != nil {
		return false, fmt.Errorf("%w - %s", err, ErrConvertDesired)
	}

	actual, err := convert.ToSliceString(actualSlice)
	if err != nil {
		return false, fmt.Errorf("%w - %s", err, ErrConvertActual)
	}

	if len(desired) != len(actual) {
		return false, nil
	}

	// ensure order
	sort.Strings(desired)
	sort.Strings(actual)

	return reflect.DeepEqual(desired, actual), nil
}

// equalSliceIntegerInterface takes in an expected slice of integers type, as an interface, converts
// it appropriately and returns the result of the comparison.
func equalSliceIntegerInterface(desiredSlice, actualSlice interface{}) (bool, error) {
	desired, err := convert.ToSliceInteger(desiredSlice)
	if err != nil {
		return false, fmt.Errorf("%w - %s", err, ErrConvertDesired)
	}

	actual, err := convert.ToSliceInteger(actualSlice)
	if err != nil {
		return false, fmt.Errorf("%w - %s", err, ErrConvertActual)
	}

	if len(desired) != len(actual) {
		return false, nil
	}

	// ensure order
	sort.Ints(desired)
	sort.Ints(actual)

	return reflect.DeepEqual(desired, actual), nil
}

// equalSliceFloatInterface takes in an expected slice of floats type, as an interface, converts
// it appropriately and returns the result of the comparison.
func equalSliceFloatInterface(desiredSlice, actualSlice interface{}) (bool, error) {
	desired, err := convert.ToSliceFloat(desiredSlice)
	if err != nil {
		return false, fmt.Errorf("%w - %s", err, ErrConvertDesired)
	}

	actual, err := convert.ToSliceFloat(actualSlice)
	if err != nil {
		return false, fmt.Errorf("%w - %s", err, ErrConvertActual)
	}

	if len(desired) != len(actual) {
		return false, nil
	}

	// ensure order
	sort.Float64s(desired)
	sort.Float64s(actual)

	return reflect.DeepEqual(desired, actual), nil
}

// equalSliceBooleanInterface takes in an expected slice of bools type, as an interface, converts
// it appropriately and returns the result of the comparison.
func equalSliceBooleanInterface(desiredSlice, actualSlice interface{}) (bool, error) {
	desired, err := convert.ToSliceBoolean(desiredSlice)
	if err != nil {
		return false, fmt.Errorf("%w - %s", err, ErrConvertDesired)
	}

	actual, err := convert.ToSliceBoolean(actualSlice)
	if err != nil {
		return false, fmt.Errorf("%w - %s", err, ErrConvertActual)
	}

	if len(desired) != len(actual) {
		return false, nil
	}

	// ensure order
	sort.Slice(desired, func(i, j int) bool {
		return convert.BooleanToInteger(desired[i]) < convert.BooleanToInteger(desired[j])
	})

	sort.Slice(desired, func(i, j int) bool {
		return convert.BooleanToInteger(actual[i]) < convert.BooleanToInteger(actual[j])
	})

	return reflect.DeepEqual(desired, actual), nil
}

// EqualSliceIntegerInterface takes in an expected slice of integers type, as an interface, converts
// it appropriately and returns the result of the comparison.
func equalSliceInterface(desiredSlice, actualSlice []interface{}) (bool, error) {
	for i := range desiredSlice {
		equal, err := Compare(desiredSlice[i], actualSlice[i])
		if !equal || err != nil {
			return false, err
		}
	}

	return true, nil
}

// equalSliceMapInterfaceInterface takes in an expected slice of map interface to interface type
// returns the result of the comparison.
func equalSliceMapInterfaceInterface(desiredSlice, actualSlice interface{}) (bool, error) {
	desired, err := convert.ToSliceMapInterfaceInterface(desiredSlice)
	if err != nil {
		return false, fmt.Errorf("%w - %s", err, ErrConvertDesired)
	}

	actual, err := convert.ToSliceMapInterfaceInterface(actualSlice)
	if err != nil {
		return false, fmt.Errorf("%w - %s", err, ErrConvertActual)
	}

	for i := range desired {
		var hasEqual bool

		for j := range actual {
			equal, err := equalMap(desired[i], actual[j])
			if err != nil {
				return false, err
			}

			if equal {
				hasEqual = true

				break
			}
		}

		if !hasEqual {
			return false, nil
		}
	}

	return true, nil
}
