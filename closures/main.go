package main

import "fmt"

func main() {
	// We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value,
	// which will be updated each time we call nextInt.
	nextInt := IntSequ()
	// See the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// To confirm that the state is unique to that particular function, create and test a new one.
	newInts := IntSequ()
	fmt.Println(newInts())
}

// IntSequ This function intSeq returns another function, which we define anonymously in the body of intSeq.
// The returned function closes over the variable i to form a closure.
func IntSequ() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
