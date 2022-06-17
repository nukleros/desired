package compare

import (
	"reflect"

	"github.com/nukleros/desired/internal/convert"
)

// EqualMap ensures any map type is equal.
func EqualMap(desired, actual interface{}) (bool, error) {
	desiredMap, err := convert.ToMapInterfaceInterface(desired)
	if err != nil {
		return false, ErrConvertDesired
	}

	actualMap, err := convert.ToMapInterfaceInterface(actual)
	if err != nil {
		return false, ErrConvertActual
	}

	for desiredKey, desiredValue := range desiredMap {
		for actualKey, actualValue := range actualMap {
			if reflect.DeepEqual(desiredKey, actualKey) {
				equal, err := Compare(desiredValue, actualValue)
				if !equal || err != nil {
					return false, err
				}
			}
		}
	}

	return true, nil
}
