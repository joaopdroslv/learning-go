package main

import (
	"fmt"
)

func main() {
	// Primitive types
	x := 42                      // int
	y := "Golang"                // string
	z := true                    // bool

	// Composite types (types built from other types)
	arr := [3]int{1, 2, 3}       // Array: Fixed-length collection of integers
	// Arrays have a fixed size, defined at declaration (size = 3 here).

	slice := []string{"A", "B"}  // Slice: Dynamic-size array-like structure
	// Slices are more flexible than arrays; they grow or shrink as needed.

	m := map[string]int{         // Map: Collection of key-value pairs
		"one": 1, 
		"two": 2,
	}
	// Maps use keys to access values. Here, "one" maps to 1, and "two" maps to 2.

	// Struct type (user-defined composite type)
	type Person struct {
		Name string  			// Name field of type string
		Age  int     			// Age field of type int
	}
	p := Person{"Alice", 30}     // Struct initialization using field values

	// Pointer type (stores the memory address of another variable)
	var a int = 10
	ptr := &a                    // &a gets the memory address of variable a
	// The pointer (ptr) now holds the address where a is stored in memory.

	// Printing values and types
	// %T prints the type
	fmt.Printf("x: Value = %d, Type = %T\n", x, x)        			// %d formats an integer
	fmt.Printf("y: Value = %s, Type = %T\n", y, y)        			// %s formats a string
	fmt.Printf("z: Value = %t, Type = %T\n", z, z)        			// %t formats a boolean
	fmt.Printf("arr: Value = %v, Type = %T\n", arr, arr)  			// %v prints the value in default format
	fmt.Printf("slice: Value = %v, Type = %T\n", slice, slice)
	fmt.Printf("map: Value = %v, Type = %T\n", m, m)
	fmt.Printf("struct (p): Value = %+v, Type = %T\n", p, p) 		// %+v includes field names in struct output
	fmt.Printf("pointer (ptr): Value = %v, Type = %T\n", ptr, ptr)
}
