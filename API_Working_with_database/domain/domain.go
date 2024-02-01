package domain

type Customer struct {
	Id          int
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      int
}

// Defining our repository Cutomer repository (This repository helps us finding all the cutomers from the server side)

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	GetCustomerById(id int) (Customer, error)
}
