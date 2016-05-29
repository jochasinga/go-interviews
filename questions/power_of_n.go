package main

import "fmt"

// Finding out if n is a power of some number can be done
// by dividing n by the number recursively until n = that number.
func isPowerOfFour(n, power uint32) bool {

	if n == 1 || n == power {
		return true
	}

	if n > 1 && n % power == 0 {
		
		n /= power

		switch {
		case n == power:
			return true
		case n > power:
			return isPowerOfFour(n, power)
		default:
			break
		}
	}
	
	return false
}
		
func main() {
	fmt.Println(isPowerOfFour(0, 4))
	fmt.Println(isPowerOfFour(1, 4))
	fmt.Println(isPowerOfFour(8, 4))
	fmt.Println(isPowerOfFour(16, 4))
	fmt.Println(isPowerOfFour(10, 2))
	fmt.Println(isPowerOfFour(64, 8))
	fmt.Println(isPowerOfFour(17, 3))
}
