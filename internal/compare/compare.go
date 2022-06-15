package compare

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

type Sorter []map[string]interface{}

func (sorter Sorter) Len() int { return len(sorter) }
func (sorter Sorter) Swap(i, j int) {
	sorter[i], sorter[j] = sorter[j], sorter[i]
}
func (sorter Sorter) Less(i, j int) bool {
	// find an equivalent key in the underlying map[string]interface{}
	for iKey, iValue := range sorter[i] {
		for jKey, jValue := range sorter[j] {
			if iKey == jKey {
				if !equalTypes(iValue, jValue) {
					continue
				}

				// prefer to find integer values and sort that way else
				// alphebetize the string values
				switch iValue.(type) {
				case int, int8, int16, int32, int64:
					return iValue.(int) < jValue.(int)
				case string:
					return !sort.StringsAreSorted(
						[]string{
							iValue.(string),
							jValue.(string),
						},
					)
				}
			}
		}
	}

	return false
}

func Compare(desiredValue, actualValue interface{}) (bool, error) {

	if !equalTypes(desiredValue, actualValue) {
		return false, errors.New(ErrorMismatchedTypes)
	}

	switch desiredAsType := desiredValue.(type) {
	case []interface{}:
		actualAsType := actualValue.([]interface{})

		// return equality if desired has no values as the
		// desired is not explicitly controlling these fields
		if len(desiredAsType) == 0 {
			return true, nil
		}

		// return inequality if actual has no values, as we have already
		// confirmed above that desired expects values
		if len(actualAsType) == 0 {
			return false, nil
		}

		switch desiredAsType[0].(type) {
		case map[string]interface{}:
			desiredSorter := Sorter(sliceInterfaceToMapStringInterface(desiredAsType))
			actualSorter := Sorter(sliceInterfaceToMapStringInterface(actualAsType))

			sort.Sort(desiredSorter)
			sort.Sort(actualSorter)

			for _, desired := range desiredSorter {
				for _, actual := range actualSorter {
					equal, err := Compare(desired, actual)
					if !equal || err != nil {
						return false, err
					}
				}
			}
		default:
			for _, desired := range desiredAsType {
				for _, actual := range actualAsType {
					equal, err := Compare(desired, actual)
					if !equal || err != nil {
						return false, err
					}
				}
			}
		}

		return true, nil
	case map[string]interface{}:
		return EqualMapStringInterface(
			desiredValue.(map[string]interface{}),
			actualValue.(map[string]interface{}),
		)
	case string:
		return EqualString(desiredValue, actualValue)
	case int, int8, int16, int32, int64:
		return EqualInt(desiredValue, actualValue)
	case bool:
		return (desiredValue.(bool) == actualValue.(bool)), nil
	case []string:
		return EqualSliceString(
			desiredValue.([]interface{}),
			actualValue.([]interface{}),
		)
	case []int, []int8, []int16, []int32, []int64:
		return EqualSliceInt(
			desiredValue.([]interface{}),
			actualValue.([]interface{}),
		)
	}

	fmt.Printf("fallthrough: %T\n\n\n", desiredValue)

	return true, nil
}

func equalTypes(desired, actual interface{}) bool {
	return reflect.TypeOf(desired) == reflect.TypeOf(actual)
}
