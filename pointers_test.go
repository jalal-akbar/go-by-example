package gobyexample

import (
	"fmt"
	"testing"
)

// We’ll show how pointers work in contrast to values with 2 functions:
// zeroval and zeroptr. zeroval has an int parameter, so arguments will be passed to it by value.
// zeroval will get a copy of ival distinct from the one in the calling function.
func zeroOval(val int) {
	val = 0
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the value at the referenced address.
func zeroPtr(iPtr *int) {
	*iPtr = 0
}

func TestPointers(t *testing.T) {
	i := 1
	fmt.Println("initial", i)

	// zeroOval
	zeroOval(i)
	fmt.Println("zerooval", i)
	// The &i syntax gives the memory address of i, i.e. a pointer to i.
	zeroPtr(&i)
	fmt.Println("zeroptr", i)
	// Pointers can be printed too.
	fmt.Println("pointer", &i)
}

// zeroval doesn’t change the i in main,
// but zeroptr does because it has a reference to the memory address for that variable.
