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

func EqualSliceMapStringInterface(desired, actual []map[string]interface{}) (bool, error) {
	for i := range desired {
		equal, err := EqualMapStringInterface(desired[i], forComparison(desired[i], actual))
		if !equal || err != nil {
			return false, err
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

func forComparison(desired map[string]interface{}, allMaps []map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	for desiredKey := range desired {
		for thisMap := range allMaps {
			for comparedKey := range allMaps[thisMap] {
				if desiredKey == comparedKey {
					return allMaps[thisMap]
				}
			}
		}
	}

	return out
}
