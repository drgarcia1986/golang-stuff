package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/drgarcia1986/golang-stuff/bubblesort"
)

func main() {
	numbers := make([]int, len(os.Args[1:]))
	for i, arg := range os.Args[1:] {
		number, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("The arg %s is invalid in position %d", arg, i)
			os.Exit(1)
		}
		numbers[i] = number
	}
	fmt.Print(bubblesort.Sort(numbers))
}
