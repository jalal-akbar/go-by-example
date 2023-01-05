package gobyexample

import (
	"fmt"
	"testing"
)

// Here we use range to sum the numbers in a slice. Arrays work like this too.

func TestRange(t *testing.T) {

	// range on arrays and slices provides both the index and value for each entry.
	// Above we didn’t need the index, so we ignored it with the blank identifier _.
	// Sometimes we actually want the indexes though.
	nums := []int{2, 3, 4}
	sum := 0

	for _, i := range nums {
		sum += i
	}
	fmt.Println(sum)

	names := []string{"jalal", "muh", "akbar"}
	for i, name := range names {
		if name != "jalal" {
			fmt.Println("index :", i)
		}
	}
	// range on map iterates over key/value pairs.
	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	for k, v := range kvs {
		fmt.Printf("Key : %s, Value %s\n", k, v)
	}
	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Printf("Key : %s\n", k)
	}
	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "ab" {
		fmt.Println(i, c)
	}
}
