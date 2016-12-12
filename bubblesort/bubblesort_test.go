package bubblesort

import "testing"

func TestSort(t *testing.T) {
	expected := []int{1, 2, 3, 5, 10}
	sorted := Sort([]int{3, 5, 1, 10, 2})
	for i := range expected {
		if expected[i] != sorted[i] {
			t.Errorf("Expected %v, got %v", expected, sorted)
		}
	}
}
