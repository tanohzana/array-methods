package arraymethods

import "testing"

func TestFilterInts(t *testing.T) {
	arrayToFilter := [5]int{1, 2, 3, 4, 5}

	filteredSlice := FilterInts(arrayToFilter[:], func(item int) bool {
		return (item > 3)
	})

	if len(filteredSlice) != 2 {
		t.Errorf("Slice size was incorrect, got: %d, want: %d.", len(filteredSlice), 2)
	}
}
