package main

import "fmt"

func countingSort(numbers []int, maxValue int) []int {

	// Allocate a numCounts slice with cap of the maxValue + 1
	// Make sure it's allocated (zero-filled) and not empty.
	numCounts := make([]int, maxValue+1)

	// For each number in numbers
	for _, num := range numbers {
		// numCounts index becomes the number, and the value
		// at each index becomes the time the number appears
		// So keeps increment when a number is found.
		numCounts[num] += 1
	}

	// Allocate the target slice
	// Make sure it's empty and not zero-filled.
	sorted := []int{}

	// Loop over the number and count in numCounts
	for num, count := range numCounts {
		// count being the time a number can appear
		// i.e. if 2 appears 3 times, 2 will be appended 3 times
		// before the above loop moves on to the next num, count
		for time := 0; time < count; time++ {
			sorted = append(sorted, num)
		}
	}

	return sorted
}

func main() {
	nums := []int{1, 2, 10, 10, 4, 6, 10, 2, 2, 54}
	sorted = countingSort(nums, 54)
	fmt.Println(sorted)
}
