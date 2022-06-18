package convert

import (
	"fmt"
	"reflect"
)

// ToMapInterfaceInterface converts an interface into a map of interface to interface.
func ToMapInterfaceInterface(in interface{}) (map[interface{}]interface{}, error) {
	switch asType := in.(type) {
	case map[interface{}]interface{}:
		return asType, nil
	default:
		if in == nil {
			return nil, ErrNilValue
		}

		asMap := reflect.ValueOf(in)

		if reflect.TypeOf(in).Kind() != reflect.Map {
			return nil, fmt.Errorf("%w - expected map but found type %T", ErrConvertMapInterfaceInterface, in)
		}

		out := make(map[interface{}]interface{})

		for _, key := range asMap.MapKeys() {
			value := asMap.MapIndex(key)

			out[key.Interface()] = value.Interface()
		}

		return out, nil
	}
}
