package compare

import (
	"reflect"
	"sort"
)

func EqualSliceInt(desired, actual []interface{}) (bool, error) {
	if len(desired) != len(actual) {
		return false, nil
	}

	desiredSliceInt, actualSliceInt := toSliceInt(desired), toSliceInt(actual)

	sort.Ints(desiredSliceInt)
	sort.Ints(actualSliceInt)

	return reflect.DeepEqual(
		desiredSliceInt,
		actualSliceInt,
	), nil
}

func toSliceInt(in []interface{}) []int {
	out := make([]int, len(in))

	for i, intValue := range in {
		out[i] = intValue.(int)
	}

	return out
}
