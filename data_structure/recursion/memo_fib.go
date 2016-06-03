package main

import "fmt"

type Fib struct {
	Memo map[uint]uint
}

func NewFib() *Fib {
	f := &Fib{
		Memo: make(map[uint]uint),
	}
	return f
}

func (f *Fib) Calculate(n uint) uint {

	// edge case
	if n < 0 {
		panic("n < 0")
	}

	// base cases
	if n == 0 || n == 1 {
		return n
	}

	for k, v := range f.Memo {
		//fmt.Printf("grabbing memo[%d]\n", n)
		if n == k {
			return v
		}
	}

	//fmt.Printf("computing fib(%d)\n", n)
	result := f.Calculate(n-1) + f.Calculate(n-2)

	// memoize
	f.Memo[n] = result

	return result

}

func main() {
	f := NewFib()
	for i := 0; i <= 10; i++ {
		fmt.Println(f.Calculate(uint(i)))
	}
}
