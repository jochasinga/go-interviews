package main

import "fmt"

type Fib struct {
	// O(1) on average
	//Memo map[uint]uint

	// O(n) worst case
	Memo []uint
}

func NewFib(max uint) *Fib {
	f := &Fib{
		//Memo: make(map[uint]uint),
		Memo: make([]uint, max),
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

	/*
		for k, v := range f.Memo {
			//fmt.Printf("grabbing memo[%d]\n", n)
			if n == k {
				return v
			}
		}
	*/

	// Map may be more time efficient
	for i, v := range f.Memo {
		//fmt.Printf("grabbing memo[%d]\n", n)
		if int(n) == i && f.Memo[uint(n)] != 0 {
			return v
		}
	}

	fmt.Printf("computing fib(%d)\n", n)
	result := f.Calculate(n-1) + f.Calculate(n-2)

	// memoize
	f.Memo[n] = result

	return result

}

func (f *Fib) InvertCalculate(n int) {
	/* Map method
	sorted := make([]uint, n+1)
	for k, v := range f.Memo {
		sorted[k] = v
	}
	*/

	// slice method
	for i := n; i >= 0; i-- {
		fmt.Println(i, f.Memo[i])
	}

}

func main() {
	f := NewFib(11)
	for i := 0; i <= 10; i++ {
		f.Calculate(uint(i))
		//fmt.Println(f.Calculate(uint(i)))
	}
	f.InvertCalculate(10)
}
