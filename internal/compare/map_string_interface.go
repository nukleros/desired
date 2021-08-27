package compare

func EqualMapStringInterface(desired, actual map[string]interface{}) (bool, error) {
	for desiredKey, desiredValue := range desired {
		for actualKey, actualValue := range actual {
			if actualKey == desiredKey {
				equal, err := Compare(desiredValue, actualValue)
				if !equal || err != nil {
					return false, err
				}
			}
		}
	}

	return true, nil
}

func sliceInterfaceToMapStringInterface(in []interface{}) []map[string]interface{} {
	out := make([]map[string]interface{}, len(in))

	for i, mapStringInterfaceValue := range in {
		out[i] = mapStringInterfaceValue.(map[string]interface{})
	}

	return out
}
