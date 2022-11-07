package gobyexample

import (
	"fmt"
	"testing"
)

// Go has various value types including strings, integers, floats, booleans, etc.
// Here are a few basic examples.

func TestValues(t *testing.T) {
	// Strings, which can be added together with +.
	fmt.Println("go" + "lang")
	// Integers
	fmt.Println("1 + 1 = ", 1+1)
	//and floats.
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// Booleans, with boolean operators as you’d expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
