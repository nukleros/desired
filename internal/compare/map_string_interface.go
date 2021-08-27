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
