// Print 1 to 100, but for every multiple of 3 print "fizz" and multiple
// of 5 print "buzz" and for multiple of both 3 and 5 print "fizzbuzz"
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("fizz")
		}
		if i%5 == 0 {
			fmt.Println("buzz")
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		}
		fmt.Println(i)
	}
}
