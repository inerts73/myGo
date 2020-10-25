package main

import "fmt"

func main() {
	fmt.Println(fib(5))
}

func fib(x int) int {
	if x <2 {
		return x
	}
	return fib(x-2) + fib(x-1)
}