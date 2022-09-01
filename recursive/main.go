package main

import "fmt"

func main() {
	fmt.Println(Fact(7))
	// Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		// Since fib was previously declared in main, Go knows which function to call with fib here.
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

// Fact This fact function calls itself until it reaches the base case of fact(0).
func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}
