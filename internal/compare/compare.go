package compare

import "fmt"

func Compare(desiredValue, actualValue interface{}) (bool, error) {

	switch desiredAsType := desiredValue.(type) {
	case []interface{}:
		// return equality if desired has no values but actual does as the
		// desired is not explicitly controlling these fields
		if len(desiredAsType) == 0 {
			return true, nil
		}

		switch desiredAsType[0].(type) {
		case map[string]interface{}:
			return EqualMapStringInterface(
				desiredValue.(map[string]interface{}),
				actualValue.(map[string]interface{}),
			)
		default:
			for _, desired := range desiredAsType {
				for _, actual := range actualValue.([]interface{}) {
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
