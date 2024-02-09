package domain

import "api/dto"

type Customer struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateofBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

// Defining our repository Cutomer repository (This repository helps us finding all the cutomers from the server side)

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	GetCustomerById(id int) (*Customer, error)
	CreateCustomer(newCustomer Customer) (*Customer, error)
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "Inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateofBirth,
		Status:      c.statusAsText(), // Calling value receiver method
	}
}

// Here in the Cutomer we are adding *(pointer) because it returns nil if the customer does'nt exists
