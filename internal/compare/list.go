package compare

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/nukleros/desired/internal/convert"
)

// EqualList takes in a list of any type and returns the result of the comparison.
func EqualList(desiredList, actualList interface{}) (bool, error) {
	desiredAsValue := reflect.ValueOf(desiredList)
	actualAsValue := reflect.ValueOf(actualList)

	// return equality if desired has no values as the
	// desired is not explicitly controlling these fields
	if desiredAsValue.Len() == 0 {
		return true, nil
	}

	// return inequality if actual has no values, as we have already
	// confirmed above that desired expects values
	if actualAsValue.Len() == 0 {
		return false, nil
	}

	// handle a map type first since ordering is important and potentially unpredictable with maps
	if desiredAsValue.Index(0).Kind() == reflect.Map {
		return EqualSliceMapInterfaceInterface(desiredList, actualList)
	}

	switch desiredAsType := desiredList.(type) {
	case []string:
		return EqualSliceStringInterface(desiredList, actualList)
	case []int, []int8, []int16, []int32, []int64:
		return EqualSliceIntegerInterface(desiredList, actualList)
	case []float32, []float64:
		return EqualSliceFloatInterface(desiredList, actualList)
	case []bool:
		return EqualSliceBooleanInterface(desiredList, actualList)
	case []interface{}:
		return EqualSliceInterface(desiredAsType, actualList.([]interface{}))
	}

	return false, fmt.Errorf("%w of type %T", ErrInvalidListType, desiredList)
}

// EqualSliceStringInterface takes in an expected slice of string type, as an interface, converts
// it appropriately and returns the result of the comparison.
func EqualSliceStringInterface(desiredSlice, actualSlice interface{}) (bool, error) {
	desired, err := convert.ToSliceString(desiredSlice)
	if err != nil {
		return false, ErrConvertDesired
	}

	actual, err := convert.ToSliceString(actualSlice)
	if err != nil {
		return false, ErrConvertActual
	}

	if len(desired) != len(actual) {
		return false, nil
	}

	// ensure order
	sort.Strings(desired)
	sort.Strings(actual)

	return reflect.DeepEqual(desired, actual), nil
}

// EqualSliceIntegerInterface takes in an expected slice of integers type, as an interface, converts
// it appropriately and returns the result of the comparison.
func EqualSliceIntegerInterface(desiredSlice, actualSlice interface{}) (bool, error) {
	desired, err := convert.ToSliceInteger(desiredSlice)
	if err != nil {
		return false, ErrConvertDesired
	}

	actual, err := convert.ToSliceInteger(actualSlice)
	if err != nil {
		return false, ErrConvertActual
	}

	if len(desired) != len(actual) {
		return false, nil
	}

	// ensure order
	sort.Ints(desired)
	sort.Ints(actual)

	return reflect.DeepEqual(desired, actual), nil
}

// EqualSliceFloatInterface takes in an expected slice of floats type, as an interface, converts
// it appropriately and returns the result of the comparison.
func EqualSliceFloatInterface(desiredSlice, actualSlice interface{}) (bool, error) {
	desired, err := convert.ToSliceFloat(desiredSlice)
	if err != nil {
		return false, ErrConvertDesired
	}

	actual, err := convert.ToSliceFloat(actualSlice)
	if err != nil {
		return false, ErrConvertActual
	}

	if len(desired) != len(actual) {
		return false, nil
	}

	// ensure order
	sort.Float64s(desired)
	sort.Float64s(actual)

	return reflect.DeepEqual(desired, actual), nil
}

// EqualSliceBooleanInterface takes in an expected slice of bools type, as an interface, converts
// it appropriately and returns the result of the comparison.
func EqualSliceBooleanInterface(desiredSlice, actualSlice interface{}) (bool, error) {
	desired, err := convert.ToSliceBoolean(desiredSlice)
	if err != nil {
		return false, ErrConvertDesired
	}

	actual, err := convert.ToSliceBoolean(actualSlice)
	if err != nil {
		return false, ErrConvertActual
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
func EqualSliceInterface(desiredSlice, actualSlice []interface{}) (bool, error) {
	if passNilComparison := compareNil(desiredSlice, actualSlice); !passNilComparison {
		return false, nil
	}

	for i := range desiredSlice {
		for j := range actualSlice {
			equal, err := Compare(desiredSlice[i], actualSlice[j])
			if !equal || err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

// EqualSliceMapInterfaceInterface takes in an expected slice of map interface to interface type
// returns the result of the comparison.
func EqualSliceMapInterfaceInterface(desiredSlice, actualSlice interface{}) (bool, error) {
	desired, err := convert.ToSliceMapInterfaceInterface(desiredSlice)
	if err != nil {
		return false, ErrConvertDesired
	}

	actual, err := convert.ToSliceMapInterfaceInterface(actualSlice)
	if err != nil {
		return false, ErrConvertActual
	}

	for i := range desired {
		var hasEqual bool

		for j := range actual {
			equal, err := EqualMap(desired[i], actual[j])
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
