package main

import "fmt"

func main() {
	// This function uses the empty interface, which means it can accept any type.
	// However, its return type is also 'interface{}', so we lose type safety
	// and need type assertions to work with the result.
	result1 := complexExample(99, 1)
	fmt.Println(result1)

	// This function uses generics.
	// The compiler enforces that 'a' and 'b' must be of the same type (int, float64, or string),
	// and the return type is preserved, giving us full type safety.
	result2 := addWithGenerics(99, 1)
	fmt.Println(result2)

	// Now i can successfully perform operations with a int typed value
	opResult := result2 + 999
	fmt.Println(opResult)
}

// Version using generics: much cleaner and safer.
// The type parameter T is constrained to int, float64, or string.
// The function guarantees that 'a' and 'b' are of the same type and returns a result of the same type.
func addWithGenerics[T int | float64 | string](a, b T) T {
	return a + b
}

// Version using the empty interface.
// Requires verbose type assertions to handle different input types,
// and the return type is interface{}, which makes it harder to use safely.
func complexExample(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)

	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}

	aString, aIsString := a.(string)
	bString, bIsString := b.(string)

	if aIsString && bIsString {
		return aString + bString
	}

	// If none of the supported type combinations match, return nil.
	return nil
}
