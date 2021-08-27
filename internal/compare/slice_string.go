package compare

import (
	"errors"
	"reflect"
	"sort"
)

func EqualSliceString(desired, actual []interface{}) (bool, error) {
	if desired == nil || actual == nil {
		return false, errors.New(ErrorNilSliceString)
	}

	if len(desired) != len(actual) {
		return false, nil
	}

	desiredSliceString := toSliceString(desired)
	actualSliceString := toSliceString(actual)

	sort.Strings(desiredSliceString)
	sort.Strings(actualSliceString)

	return reflect.DeepEqual(
		desiredSliceString,
		actualSliceString,
	), nil
}

func toSliceString(in []interface{}) []string {
	out := make([]string, len(in))

	for i, stringValue := range in {
		out[i] = stringValue.(string)
	}

	return out
}
