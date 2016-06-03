package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {

	result := []int{}

	// O(2n) space
	num1 := map[int]int{}
	num2 := map[int]int{}

	// O(2n) time
	for _, n1 := range nums1 {
		num1[n1]++
	}

	for _, n2 := range nums2 {
		num2[n2]++
	}

	// O(n^2) time
	for n1, t1 := range num1 {
		for n2, t2 := range num2 {

			if n1 == n2 {
				if t1 >= t2 {
					for i := 0; i < t2; i++ {
						result = append(result, n1)
					}
				}
				if t1 < t2 {
					for i := 0; i < t1; i++ {
						result = append(result, n1)
					}
				}
			}
		}
	}
	return result
}

/*
func intersectWithBinarySearch(nums1 []int, nums2 []int) []int {

	result := []int{}

	for _, n2 := range nums2 {

		target := n2

		// binary search
		floorIndex := -1
		ceilIndex := len(nums1)

		for floorIndex + 1 < ceilIndex {
			dist := ceilIndex - floorIndex
			var halfDist int = dist / 2
			guessIndex := floorIndex + halfDist

			guess := nums1[guessIndex]

			switch {
			case guess == target:
				result = append(result, guess)
				fallthrough
			case guess > target:
				ceilIndex = guessIndex
				fallthrough
			default:
				floorIndex = guessIndex
			}
		}
	}
	return result
}
*/

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}
