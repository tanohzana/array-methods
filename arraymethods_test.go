package arraymethods

import (
	"strconv"
	"testing"
)

func TestFilterInts(t *testing.T) {
	arrayToFilter := [5]int{1, 2, 3, 4, 5}

	filteredSlice := FilterInts(arrayToFilter[:], func(item int) bool {
		return (item > 3)
	})

	if len(filteredSlice) != 2 {
		t.Errorf("Slice size was incorrect, got: %d, want: %d... :-(", len(filteredSlice), 2)
	}
}

func TestFilterStrings(t *testing.T) {
	arrayToFilter := [5]string{"cat", "dog", "horse", "dinosaur", "bird"}

	filteredSlice := FilterStrings(arrayToFilter[:], func(item string) bool {
		return (string(item[0]) == "d")
	})

	if len(filteredSlice) != 2 {
		t.Errorf("Slice size was incorrect, got: %d, want: %d... :-(", len(filteredSlice), 2)
	}
}

func TestFilterJson(t *testing.T) {
	expectedLength := 1
	jsonContent1 := LoadJSONAsMap("./test_utils/jsonTestFile1.json")
	jsonContent2 := LoadJSONAsMap("./test_utils/jsonTestFile2.json")
	jsonContent3 := LoadJSONAsMap("./test_utils/jsonTestFile3.json")

	jsonArray := [3]map[string]interface{}{jsonContent1, jsonContent2, jsonContent3}

	filteredJSONArray := FilterJSON(jsonArray[:], func(item map[string]interface{}) bool {
		returnValue := true

		for _, value := range item {
			if isTrue, err := strconv.ParseBool(value.(string)); err == nil && !isTrue {
				returnValue = false
			}
		}

		return returnValue
	})

	if len(filteredJSONArray) != expectedLength {
		t.Errorf("Expected %d, but got: %d", expectedLength, len(filteredJSONArray))
	}
}

func TestMapInts(t *testing.T) {
	expectedArray := [5]int{2, 4, 6, 8, 10}
	intsArray := [5]int{1, 2, 3, 4, 5}

	mappedArray := MapInts(intsArray[:], func(entry int) int {
		return 2 * entry
	})

	for i := 0; i < len(expectedArray); i++ {
		if expectedArray[i] != mappedArray[i] {
			t.Errorf("Expected %d, but got %d, in position %d", expectedArray[i], mappedArray[i], i)
		}
	}
}
