package main

import "fmt"

func main() {
	age := 32 // Regular variable

	var agePointer *int // creating the pointer variable 

	agePointer = &age //  address of the age is given to the poniter variable
	

	fmt.Println("Age: ", *agePointer) // Now this line of code will give the address of the variable
	//the *pointer dereference the pointer
	editAgetoadultyears(agePointer)
	fmt.Println("The edited age is :", age) // here we are o/p the age which which is changed
}

func editAgetoadultyears(age *int) {
	*age = *age - 18
}

// Here We are directly overrinding the gievn age value int adult years by using the function editAgetoadultyears
