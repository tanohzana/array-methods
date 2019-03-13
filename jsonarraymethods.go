package jsonmapmethods

func Keys(jsonMap map[string]interface{}) []string {
	var keysArray []string

	for key, _ := range jsonMap {
		keysArray = append(keysArray, key)
	}

	return keysArray
}

func Values(jsonMap map[string]interface{}) []interface{} {
	var valuesArray []interface{}

	for _, value := range jsonMap {
		valuesArray = append(valuesArray, value)
	}

	return valuesArray
}
