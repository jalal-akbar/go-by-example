package gobyexample

import (
	"fmt"
	"testing"
)

// The (int, int) in this function signature shows that the function returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

func TestMultipeReturn(t *testing.T) {
	// Here we use the 2 different return values from the call with multiple assignment.
	a, b := vals()
	fmt.Printf("Value a = %v\nValue b = %v\n", a, b)
	// If you only want a subset of the returned values, use the blank identifier _.
	_, c := vals()
	fmt.Printf("Value c = %v\n", c)
}
