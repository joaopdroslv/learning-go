package main

import (
	"fmt"
	"math"
)

// Constants can be defined outside a function, here it has file scope
const inflationRate = 2.5 // Can't be reasigned

func main() {
	var investmentAmount float64 // Defining the type explicity
	var years float64

	// Setting a default value
	expectedReturnRate := 3.31 // Letting Go define implicitly

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount) // Pass the variable's pointer so fmt.Scan can store the user input

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// Use short variable declaration; Go infers the type automatically
	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, years, expectedReturnRate)

	// Easier way to print
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Real Value (ajusted for inflation):", futureRealValue)

	// fmt.Sprintf formats the output and stores it in a string for later use
	// formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	// formattedFRV := fmt.Sprintf("Future Real Value (ajusted for inflation): %.2f\n", futureRealValue)
	// fmt.Print(formattedFV, formattedFRV)

	// Using fmt.Printf to format output with two decimal places
	// %.2f formats a float with 2 digits after the decimal point
	fmt.Printf("Future Value: %.2f\nFuture Real Value (ajusted for inflation): %.2f\n", futureValue, futureRealValue)

	// Using backticks
	// fmt.Printf(`Future Value: %.2f
	// Future Real Value (adjusted by inflation): %.2f`, futureValue, futureRealValue)
}

// Defining a simple function
func outputText(text string) {
	fmt.Print(text)
}

// Defining a function that accepts multiple parameters and returns multiple values
// func calculateFutureValues(investmentAmount, years, expectedReturnRate float64) (float64, float64) {
// 	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
// 	frv := fv / math.Pow((1+inflationRate/100), years)
// 	return fv, frv
// }

// Different approach
func calculateFutureValues(investmentAmount, years, expectedReturnRate float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv = fv / math.Pow((1+inflationRate/100), years)
	return
}
