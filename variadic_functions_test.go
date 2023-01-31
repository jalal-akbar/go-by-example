package gobyexample

import (
	"fmt"
	"testing"
)

func sum(nums ...int) {
	// Here’s a function that will take an arbitrary number of ints as arguments.
	fmt.Print(nums, " ")
	total := 0
	//Within the function, the type of nums is equivalent to []int.
	//We can call len(nums), iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func TestVariadicFunctions(t *testing.T) {
	// Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)
	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...) like this.
	num := []int{1, 2, 3}
	sum(num...)
}
