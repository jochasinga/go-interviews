package main

import "fmt"

// Finding out if n is a power of some number can be done
// by dividing n by the number recursively until n = that number.
func isPowerOfFour(n uint32) bool {

	n /= 4

	switch {
	case n == 4:
		return true
	case n > 4:
		return isPowerOfFour(num)
	default:
		break
	}
	
	return false
}
		
func main() {
	fmt.Println(isPowerOfFour(8))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(10))
	fmt.Println(isPowerOfFour(64))
}
