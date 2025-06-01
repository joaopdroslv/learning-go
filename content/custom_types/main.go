package main

import "fmt"

// Defining a custom type based on the built-in string type
type str string

// Method attached to the custom type `str`
// Allows any `str` value to call `.log()` and print itself
func (text str) log() {
	fmt.Println(text)
}

func main() {
	// Declare a variable of type `str`
	var name str = "John"

	// Call the custom method defined for the `str` type
	name.log()
}
