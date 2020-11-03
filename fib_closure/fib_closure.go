package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) []int {
	//var fib []int = []int{}
	fib := make([]int,0)
	
	return func(N int) []int {		
		if N < 2 {
			fib = append(fib, N)
			return fib
		}
		
		fib = append(fib, fib[N-2] + fib[N-1])
		
		return fib
	}
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
