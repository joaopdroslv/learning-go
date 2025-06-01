package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstNameInput := getUserData("Please enter your first name: ")
	lastNameInput := getUserData("Please enter your last name: ")
	birthDateInput := getUserData("Please enter your birthdate (mm/dd/yyyy): ")

	var appUser *user.User

	// Using a contructor to create a new user
	appUser, err := user.New(firstNameInput, lastNameInput, birthDateInput)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "admin")

	admin.OutputUserDetails()

	// Call the method attached to the user struct
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
