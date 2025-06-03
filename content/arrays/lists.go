package main

import "fmt"

func main() {
	var names [4]string // Declare an empty array of 4 strings
	fmt.Println(names)
	// >> [   ] => an array with 4 empty string slots (default zero values)

	// Initialize the first element of the array
	names = [4]string{"Book"}
	fmt.Println(names)
	// >> [Book   ] => array with only the first element set, rest are empty

	// Array declaration syntax: [SIZE]TYPE{VALUES}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	// >> [10.99 9.99 45.99 20]

	// ---------- Accessing values

	// Access an individual element by index
	fmt.Println(names[0])
	// >> Book

	// Set a value at a specific index
	names[3] = "Magazine"
	fmt.Println(names)
	// >> [Book   Magazine] => elements at indexes 0 and 3 are set

	// ---------- Slicing arrays

	// Create a slice from index 1 to 3 (3 is excluded)
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
	// >> [9.99 45.99]

	// Slice from start to index 3 (excluding 3)
	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)
	// >> [10.99 9.99 45.99]

	// Slice from index 1 to the end
	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)
	// >> [9.99 45.99 20]

	// Slices are views into the original array.
	// Modifying a slice also modifies the underlying array.
	// You're not copying data — just referencing a part of it.
	fmt.Println(prices)
	// >> [10.99 9.99 45.99 20]
	featuredPrices[0] = 11.99
	fmt.Println(prices)
	// >> [10.99 11.99 45.99 20] => change is reflected in the original array

	// ---------- Built-in functions

	fmt.Println(len(prices), cap(prices)) // Get length and capacity of the array
	// >> 4 4

	// Slice a slice: keep only the first item
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	// >> [11.99] | 1 3

	// Slices have a length and a capacity.
	// Length: number of elements currently in the slice.
	// Capacity: how many elements can be added (until it reaches the end of the underlying array).
	// You can extend a slice up to its capacity.
	fmt.Println(highlightedPrices[0])

	// Extend the slice to include more elements
	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	// >> [11.99 45.99 20] | 3 3
}
