package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// Creating a new instance of User
	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("meghana@3467", "kavyau^78")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
	// Output user details
	appUser.OutputUserDetails()

	// Clear user name
	appUser.ClearUserName()

	// Output user details again
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
