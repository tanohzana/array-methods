package arraymethods

func FilterInts(objToFilter []int, f func(int) bool) []int {
	var filteredSlice []int

	for i := 0; i < len(objToFilter); i++ {
		if f(objToFilter[i]) {
			filteredSlice = append(filteredSlice, objToFilter[i])
		}
	}

	return filteredSlice
}

func Map(objectToMap []interface{}, mapFunction func(interface{}) interface{}) []interface{} {
	var mappedObj []interface{}

	for key, value := range objectToMap {
		mappedObj[key] = mapFunction(value)
	}

	return mappedObj
}
