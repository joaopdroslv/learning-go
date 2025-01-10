package main

import (
	"fmt"
)

func main() {
	// Operators and Expressions in Go are the same as many other languages

	// 1. Arithmetic
	fmt.Println("Arithmetic:")
	
	a, b := 10, 5
	
	fmt.Println("a + b =", a + b)  // +: Addition
	fmt.Println("a - b =", a - b)  // -: Subtraction
	fmt.Println("a * b =", a * b)  // *: Multiplication
	fmt.Println("a / b =", a / b)  // /: Division
	fmt.Println("a % b =", a % b)  // %: Modulus

	// 2. Relational
	fmt.Println("\nRelational:")

	fmt.Println("a == b:", a == b)  // is equal
	fmt.Println("a != b:", a != b)  // is not equal
	fmt.Println("a > b:", a > b)    // is greater than
	fmt.Println("a < b:", a < b)    // is less than
	fmt.Println("a >= b:", a >= b)  // is greater than or equal
	fmt.Println("a <= b:", a <= b)  // is less than or equal

	// 3. Logical
	fmt.Println("\nLogical:")

	c, d := true, false

	fmt.Println("c && d:", c && d)  // AND
	fmt.Println("c || d:", c || d)  // OR
	fmt.Println("!c:", !c)          // NOT
}
