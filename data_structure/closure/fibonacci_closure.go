package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int, so on and so forth.
// fib: prev, recent, prev + recent, ...
// fib(0) = 0
// fib(1) = 1
func fibonacci() func(int) int {
	// Start from 0
	prev := 0
	// Then recent 1
	recent := 1
	// Current is 1
	current := prev + recent
	return func(n int) int {
		var ret int
		switch n {
		// Handle each case
		case 0:
			n++
			ret = 0
		case 1:
			n++
			ret = 1
		default:
			ret = current
			prev = recent
			recent = current
			current = prev + recent
		}
		return ret
	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}

}
