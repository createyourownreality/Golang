package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct { // Here the struct and the type of the struct is created
	firstName string
	lastName  string
	birthDate string
	Age       int
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     // 		this user struct is embbeded or reused
}

func (u User) OutputUserDetails() { // here we are passing the parameter as u and the type as user
	fmt.Println(u.firstName, u.lastName, u.birthDate) // Here we using the "." operator and accessing the struct members
}

// now the above block of code is method and it is not a function { So inside the method the "u" is the instance of the User }

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email, // here we are calling the required Admin block
		password: password,
		User: User{
			firstName: "Admin", // here we are calling the embeded struct
			lastName:  "Admin",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func NewUser(userfirstName, userlastName, userbirthdate string) (*User, error) {
	if userfirstName == "" || userlastName == "" || userbirthdate == "" {
		return nil, errors.New("Firstname, lastname and birthname are required. ")
	}
	return &User{
		firstName: userfirstName,
		lastName:  userlastName,
		birthDate: userbirthdate, // Here we  are creating an instance for the User struct and the we get the instance value
		createdAt: time.Now(),
	}, nil
}

// Here (u User is the Reciver , indicating that this method is asscoiated with this "User" type)
// u is the name of the
