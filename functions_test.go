package gobyexample

import (
	"fmt"
	"testing"
)

// Here’s a function that takes two ints and returns their sum as an int.

// Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.

func plus(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters up to the final parameter that declares the type.

func plusPlus(a, b, c int) int {
	return a + b + c
}

// Call a function just as you’d expect, with name(args).
func TestFunctions(t *testing.T) {
	var a, b, c int = 2, 3, 4
	fmt.Println("Plus", plus(a, b))

	fmt.Println("PlusPlus", plusPlus(a, b, c))
}
