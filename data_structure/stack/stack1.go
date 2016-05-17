package main

import "fmt"

type Stack []int

func (s Stack) push(n int) Stack {
	s = append(s, n)
	return s
}

func (s Stack) pop() (Stack, int) {
	popped := s[len(s)-1]
	s = s[:len(s)-1]
	return s, popped
}

func main() {
	s := Stack{}
	for i := 1; i <= 4; i++ {
		s = s.push(i)
	}
	fmt.Println(s)
	s, n := s.pop()
	fmt.Println(s, n)
}

