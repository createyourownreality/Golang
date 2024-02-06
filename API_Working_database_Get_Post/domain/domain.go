package domain

type Customer struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"DateofBirth"`
	Status      int    `json:"Status"`
}

// Defining our repository Cutomer repository (This repository helps us finding all the cutomers from the server side)

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	GetCustomerById(id int) (*Customer, error)
	CreateCustomer(newCustomer Customer) (*Customer, error)
}

// Here in the Cutomer we are adding *(pointer) because it returns nil if the customer does'nt exists
