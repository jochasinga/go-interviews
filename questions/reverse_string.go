package main

import "fmt"

func reverseString(s string) string {
	if len(s) == 0 {
		return s
	}

	n := 0
	target := make([]rune, len(s))

	for _, char := range s {
		target[n] = char
		n++
	}
	//target = target[0:n]

	for i := 0; i < n/2; i++ {
		target[i], target[n-1-i] = target[n-1-i], target[i]
	}

	return string(target)
}

func main() {
	fmt.Println(reverseString("hello"))
}
