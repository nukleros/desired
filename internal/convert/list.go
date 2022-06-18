package convert

import (
	"fmt"
	"reflect"
)

// ToSliceString converts an interface to a slice of strings and returns an error
// if unable to.
func ToSliceString(in interface{}) ([]string, error) {
	if in == nil {
		return nil, ErrNilValue
	}

	// attempt direct conversion first
	outDirect, ok := in.([]string)
	if ok {
		return outDirect, nil
	}

	// attempt indirect conversion next
	outIndirect, ok := in.([]interface{})
	if !ok {
		return nil, ErrConvertSliceString
	}

	out := make([]string, len(outIndirect))

	for i := range outIndirect {
		outStr, err := ToString(outIndirect[i])
		if err != nil {
			return nil, ErrConvertSliceString
		}

		out[i] = outStr
	}

	return out, nil
}

// ToSliceInteger converts an interface to a slice of integers and returns an error
// if unable to.
func ToSliceInteger(in interface{}) ([]int, error) {
	if in == nil {
		return nil, ErrNilValue
	}

	// attempt direct conversion first
	outDirect, ok := in.([]int)
	if ok {
		return outDirect, nil
	}

	// attempt indirect conversion next
	outIndirect, ok := in.([]interface{})
	if !ok {
		return nil, ErrConvertSliceInteger
	}

	out := make([]int, len(outIndirect))

	for i := range outIndirect {
		outInt, err := ToInteger(outIndirect[i])
		if err != nil {
			return nil, ErrConvertSliceInteger
		}

		out[i] = outInt
	}

	return out, nil
}

// ToSliceFloat converts an interface to a slice of float64s and returns an error
// if unable to.
func ToSliceFloat(in interface{}) ([]float64, error) {
	if in == nil {
		return nil, ErrNilValue
	}

	// attempt direct conversion first
	outDirect, ok := in.([]float64)
	if ok {
		return outDirect, nil
	}

	// attempt indirect conversion next
	outIndirect, ok := in.([]interface{})
	if !ok {
		return nil, ErrConvertSliceFloat
	}

	out := make([]float64, len(outIndirect))

	for i := range outIndirect {
		outFloat, err := ToFloat(outIndirect[i])
		if err != nil {
			return nil, ErrConvertSliceFloat
		}

		out[i] = outFloat
	}

	return out, nil
}

// ToSliceBoolean converts an interface to a slice of strings and returns an error
// if unable to.
func ToSliceBoolean(in interface{}) ([]bool, error) {
	if in == nil {
		return nil, ErrNilValue
	}

	// attempt direct conversion first
	outDirect, ok := in.([]bool)
	if ok {
		return outDirect, nil
	}

	// attempt indirect conversion next
	outIndirect, ok := in.([]interface{})
	if !ok {
		return nil, ErrConvertSliceString
	}

	out := make([]bool, len(outIndirect))

	for i := range outIndirect {
		outBool, err := ToBoolean(outIndirect[i])
		if err != nil {
			return nil, ErrConvertSliceString
		}

		out[i] = outBool
	}

	return out, nil
}

// ToSliceMapInterfaceInterface converts an interface to a slice of map interface to interface and returns an error
// if unable to.
func ToSliceMapInterfaceInterface(in interface{}) ([]map[interface{}]interface{}, error) {
	if in == nil {
		return nil, ErrNilValue
	}

	asList := reflect.ValueOf(in)

	if reflect.TypeOf(in).Kind() != reflect.Slice && reflect.TypeOf(in).Kind() != reflect.Array {
		return nil, fmt.Errorf("%w - expected list but found type %T", ErrConvertSliceMapInterfaceInterface, in)
	}

	out := make([]map[interface{}]interface{}, asList.Len())

	for i := 0; i < asList.Len(); i++ {
		asMap, err := ToMapInterfaceInterface(asList.Index(i).Interface())
		if err != nil {
			return nil, ErrConvertSliceMapInterfaceInterface
		}

		out[i] = asMap
	}

	return out, nil
}
