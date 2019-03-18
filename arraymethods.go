package arraymethods

// FilterInts allows to loop on int slices and filter them according to the function passed in argument
func FilterInts(objToFilter []int, f func(int) bool) []int {
	var filteredSlice []int

	for i := 0; i < len(objToFilter); i++ {
		if f(objToFilter[i]) {
			filteredSlice = append(filteredSlice, objToFilter[i])
		}
	}

	return filteredSlice
}

// FilterStrings allows to loop on string slices and filter them according to the function passed in argument
func FilterStrings(sliceToFilter []string, filterFunc func(string) bool) []string {
	var filteredSlice []string

	for i := 0; i < len(sliceToFilter); i++ {
		if filterFunc(sliceToFilter[i]) {
			filteredSlice = append(filteredSlice, sliceToFilter[i])
		}
	}

	return filteredSlice
}

// FilterJSON allows to loop on JSON (map[string]interface{}) slices and filter them according to the function passed in argument
func FilterJSON(sliceToFilter []map[string]interface{}, filterFunc func(item map[string]interface{}) bool) []map[string]interface{} {
	var filteredArray []map[string]interface{}

	for i := 0; i < len(sliceToFilter); i++ {
		if filterFunc(sliceToFilter[i]) {
			filteredArray = append(filteredArray, sliceToFilter[i])
		}
	}

	return filteredArray
}

func MapInts(sliceToMap []int, mappingFunc func(int) int) []int {
	var mappedInts []int

	for i := 0; i < len(sliceToMap); i++ {
		mappedInts = append(mappedInts, mappingFunc(sliceToMap[i]))
	}

	return mappedInts
}
