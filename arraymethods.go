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

func FilterStrings(sliceToFilter []string, filterFunc func(string) bool) []string {
	var filteredSlice []string

	for i := 0; i < len(sliceToFilter); i++ {
		if filterFunc(sliceToFilter[i]) {
			filteredSlice = append(filteredSlice, sliceToFilter[i])
		}
	}

	return filteredSlice
}

func FilterJSON(sliceToFilter []map[string]interface{}, filterFunc func(item map[string]interface{}) bool) []map[string]interface{} {
	var filteredArray []map[string]interface{}

	for i := 0; i < len(sliceToFilter); i++ {
		if filterFunc(sliceToFilter[i]) {
			filteredArray = append(filteredArray, sliceToFilter[i])
		}
	}

	return filteredArray
}
