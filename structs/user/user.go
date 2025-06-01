package user

import (
	"errors"
	"fmt"
	"time"
)

// user defines the structure for storing user-related data
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Embedding a User struct inside Admin gives Admin access to User's methods.
// You can embed it with a field name (e.g., "User User") or anonymously (just "User").
// Named embedding (User User) requires accessing methods like admin.User.MethodName().
// Anonymous embedding (User) allows calling methods directly: admin.MethodName().
type Admin struct {
	email    string
	password string
	User
}

// This syntax attaches the function to the user struct,
// making it a method with receiver u of type user (non-pointer).
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

// This method demonstrates how to modify struct data.
// Always use a pointer receiver when you need to change the original struct;
// otherwise, you'll be working with a copy and changes won't persist.
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Using a contructor to set validation logic
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	// Could add validation here too

	return Admin{
		email:    email,
		password: password,
		// If using named embedding (e.g., user User), use the field name explicitly:
		// user: User{...}
		// If using anonymous embedding (just User), it's the same syntax,
		// but method access becomes simpler: admin.MethodName() instead of admin.User.MethodName()
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "--/--/----",
		},
	}
}

// Standalone function to output user details from a pointer receiver
// func outputUserDetails(u *user) {
// 	// While (*u).firstName is the explicit way to dereference the pointer,
// 	// Go allows accessing struct fields directly through the pointer using u.firstName.
// 	// The compiler automatically handles the dereferencing.
// 	fmt.Println((*u).firstName, u.lastName, u.birthDate, u.createdAt)
// }
