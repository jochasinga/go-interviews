// A general algorithm to reverse an array/slice
package main

import "fmt"

// Using recursion
func reverseInt1(input []int) []int {
	if len(input) == 0 {
		return input
	}

	return append(reverseInt1(input[1:]), input[0])
}

// Using in-place swapping
func reverseInt2(input []int) []int {
	if len(input) == 0 {
		return input
	}

	n := len(input)
	for i := 0; i < n/2; i++ {
		input[i], input[n-1-i] = input[n-1-i], input[i]
	}

	return input
}

func main() {
	nums := []int{1, 2, 3, 5, 8}
	fmt.Println(reverseInt2(nums))
}
