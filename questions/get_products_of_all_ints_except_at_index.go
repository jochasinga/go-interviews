package main

import "fmt"

func getProductsOfAllIntsExceptAtIndex(nums []int) []int {

	// BASE CASE: Make sure there is more than one member
	if len(nums) <= 1 {
		panic("Slice have less than two members!")
	}

	// Allocate a new variable O(1) space
	new := make([]int, len(nums))

	// This is O(n) time
	for i := range nums {
		switch i {
		case 0:
			new[i] = reduceMultiply(nums[i+1:])
		case len(nums) - 1:
			new[i] = reduceMultiply(nums[:i])
		default:
			// This is O(2) space
			first := nums[:i]
			second := nums[i+1:]
			first = append(first, second...)
			new[i] = reduceMultiply(first)
		}
	}
	return new
}

func reduceMultiply(nums []int) int {
	product := nums[0]
	for i := range nums {
		if i == 0 {
			continue
		}
		fmt.Printf("round %d: %d\n", i, product)
		product *= nums[i]

	}
	return product
}

func main() {
	a := []int{1, 7, 3, 4}
	//b := []int{1, 7, 3, 0}
	fmt.Println(getProductsOfAllIntsExceptAtIndex(a))
	//fmt.Println(getProductsOfAllIntsExceptAtIndex(b))
}
