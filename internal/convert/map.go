package convert

import (
	"fmt"
	"reflect"
)

// ToMapInterfaceInterface converts an interface into a map of interface to interface.
func ToMapInterfaceInterface(in interface{}) (map[interface{}]interface{}, error) {
	if in == nil {
		return nil, ErrNilValue
	}

	asMap := reflect.ValueOf(in)

	if asMap.Kind() != reflect.Map {
		return nil, fmt.Errorf("%w - expected map but found type %T", ErrConvertMapInterfaceInterface, in)
	}

	out := make(map[interface{}]interface{})

	for _, key := range asMap.MapKeys() {
		value := asMap.MapIndex(key)

		out[key.Interface()] = value.Interface()
	}

	return out, nil
}
