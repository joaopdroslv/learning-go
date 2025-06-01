package main

import (
	"fmt"
	"time"
)

type user struct { // Define our structure
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	firstNameInput := getUserData("Please enter your first name: ")
	lastNameInput := getUserData("Please enter your last name: ")
	birthDateInput := getUserData("Please enter your birthdate (mm/dd/yyyy): ")

	// fmt.Println(firstNameInput, lastNameInput, birthDateInput)

	var appUser user

	appUser = user{
		firstName: firstNameInput,
		lastName:  lastNameInput,
		birthDate: birthDateInput,
		createdAt: time.Now(),
	}

	outputUserDetails(&appUser)

	// ... do something awesome with that gathered data!
}

func outputUserDetails(u *user) {
	// While (*u).firstName is the explicit way to dereference the pointer,
	// Go allows accessing struct fields directly through the pointer using u.firstName.
	// The compiler automatically dereferences the pointer
	fmt.Println((*u).firstName, u.lastName, u.birthDate, u.createdAt)
}

func getUserData(prompText string) string {
	fmt.Print(prompText)
	var value string
	fmt.Scan(&value)
	return value
}
