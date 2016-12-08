package bubblesort_test

import (
	"testing"

	"github.com/drgarcia1986/golang-stuff/bubblesort"
)

func TestSort(t *testing.T) {
	expected := []int{1, 2, 3, 5, 10}
	sorted := bubblesort.Sort([]int{3, 5, 1, 10, 2})
	for i := range expected {
		if expected[i] != sorted[i] {
			t.Errorf("Expected %v, got %v", expected, sorted)
		}
	}
}
