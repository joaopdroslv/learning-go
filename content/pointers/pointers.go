package main

import "fmt"

func main() {
	age := 32 // Regular integer variable

	var agePointer *int
	agePointer = &age // Use & to get the memory address (pointer) of the variable

	fmt.Println("Age:", *agePointer) // Use * to dereference the pointer and access its value

	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears) // New value calculated, but original 'age' remains unchanged
	fmt.Println(age)        // 'age' is still 32

	editAgeToAdultYears(agePointer)
	fmt.Println(age) // 'age' is now updated with the new value from inside the function
}

// Instead of passing the value directly (which would create a copy),
// we pass a pointer to the variable's memory address.
// This allows the function to read the original value without creating a copy.

func getAdultYears(age *int) int { // *int indicates the function expects a pointer to an int
	return *age - 18 // Dereference the pointer to access the value (without modifying it)
}

func editAgeToAdultYears(age *int) { // *int indicates the function expects a pointer to an int
	*age = *age - 18 // Dereference and assign a new value (modifies the original variable)
}
