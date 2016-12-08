package bubblesort

func Sort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	pivotIdx := len(numbers) / 2
	pivot := numbers[pivotIdx]

	numbersWithoutPivot := append(numbers[:pivotIdx], numbers[pivotIdx+1:]...)

	minors, biggers := separate(numbersWithoutPivot, pivot)
	return append(append(Sort(minors), pivot), Sort(biggers)...)

}

func separate(numbers []int, pivot int) (minors, biggers []int) {
	for _, n := range numbers {
		if n <= pivot {
			minors = append(minors, n)
		} else {
			biggers = append(biggers, n)
		}
	}
	return
}
