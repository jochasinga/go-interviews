func intersection(nums1 []int, nums2 []int) []int {
	// This is a O(n^2) case though...

	result := []int{}
	isExist := map[int]bool{}

	// loop the first list
	for _, n1 := range nums1 {

		// for each value, see if it matches the second list
		for _, n2 := range nums2 {

			// if found ...
			if n1 == n2 {

				// if this is the first hit...
				if len(isExist) == 0 {

					// Append to result right away
					result = append(result, n2)

					// Record the value in isExist map
					isExist[n2] = true
				}

				// Now that it's not first hit,
				// let's check if the next hit is already in isExist
				if isExist[n2] == false {
					result = append(result, n2)
					isExist[n2] = !isExist[n2]
				}

			}
		}
	}

	return result
}
