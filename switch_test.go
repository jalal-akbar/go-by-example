package gobyexample

import (
	"fmt"
	"testing"
	"time"
)

// Switch statements express conditionals across many branches

func TestSwitch(t *testing.T) {
	// Here’s a basic switch.
	i := 2
	fmt.Print("Write ", i, " as")
	switch i {
	case 1:
		fmt.Println(" One")
	case 2:
		fmt.Println(" Two")
	case 3:
		fmt.Println(" Three")
	}
	// You can use commas to separate multiple expressions in the same case statement.
	//We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is a Weekend")
	default:
		fmt.Println("It is a Weekday")
	}
	// switch without an expression is an alternate way to express if/else logic.
	// Here we also show how the case expressions can be non-constants.
	w := time.Now()
	switch {
	case w.Hour() < 12:
		fmt.Println("Before noon")
	default:
		fmt.Println("After noon")
	}
	// A type switch compares types instead of values.
	// You can use this to discover the type of an interface value.
	// In this example,
	// the variable t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch typ := i.(type) {
		case bool:
			fmt.Println("i am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("Dont Know Type %T\n", typ)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}
