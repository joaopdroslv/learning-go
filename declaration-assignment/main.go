package main

import (
	"fmt"
)

// The ":=" short declaration operator cannot be used here; it only works inside function blocks.
var y = "Good Morning!"  // Package-level variable (available globally within the package)

func main() {
	// Variable Declaration and Assignment
	// Variables declared inside a function are only available within the function's scope.

	// Declaration and Initialization
	x := 10  // ":=" declares and initializes a variable with type inferred from the assigned value.
	// In this case, x is inferred to be of type int.

	fmt.Printf("x: %v, %T\n", x, x)  // %v prints the value, %T prints the type

	// Assignment (reassigning a new value to x)
	x = 20  // The type of x remains int, as it was inferred during the first declaration.

	fmt.Printf("x: %v, %T\n", x, x)
}
