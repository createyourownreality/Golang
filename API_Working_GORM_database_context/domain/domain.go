package domain

import (
	"api/dto"
	"context"
	"time"
)

type Customer struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	Name        string    ` json:"name"`
	City        string    `json:"city"`
	Zipcode     string    ` json:"zipcode"`
	DateofBirth time.Time ` gorm:"column:date_of_birth" json:"date_of_birth"`
	Status      string    ` json:"status"`
}

// Defining our repository Cutomer repository (This repository helps us finding all the cutomers from the server side)
// this sits in the boundary of the domain (this the customerrepository interface)
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	GetCustomerById(ctx context.Context, id int) (*Customer, error)
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
