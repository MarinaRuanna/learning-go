package main

import "fmt"

func main() {
	Plus(1, 3)
	fmt.Println(PlusPlus(1, 2, 3))

	a, b := Vals()
	fmt.Println(a)
	fmt.Println(b)
	// PS: If you only want a subset of the returned values, use the blank identifier _.
	_, c := Vals()
	fmt.Println(c)

	Sum(1, 2)
	Sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	Sum(nums...)

}

func Plus(a int, b int) int {
	return a + b
}

// PlusPlus The (int, int) in this function signature shows that the function returns 2 ints.
func PlusPlus(a, b, c int) int {
	return a + b + c
}

// Vals Here we use the 2 different return values from the call with multiple assignment.
func Vals() (int, int) {
	return 3, 7
}

// Sum Here’s a function that will take an arbitrary number of ints as arguments. Within the function,
// the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.
func Sum(nums ...int) {

	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}
