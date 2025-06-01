package main

import (
	"fmt"

	"example.com/bank/file_utils"
	"github.com/Pallinder/go-randomdata" // Using third-party libraries
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = file_utils.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("[ERROR] Something went wrong!")
		fmt.Println(err)
		fmt.Println("----------")
		// panic("Can't continue without balance information.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)

		case 2:
			fmt.Print("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount

		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. Insufficient founds!")
				continue
			}

			accountBalance -= withdrawAmount

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}

		if choice == 2 || choice == 3 {
			fmt.Println("Balance updated! New balance:", accountBalance)
			file_utils.WriteFloatToFile(accountBalance, accountBalanceFile)
		}
	}
}
