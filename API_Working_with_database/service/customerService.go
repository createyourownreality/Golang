package service

import (
	"api/domain"
)

// Here we are getting all the customers
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(id int) (domain.Customer, error)
}

// Here we are creating the implementation

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// reciver function an attaching the defaultCustomerService
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// helper function to instatiate(This used for creating the instances of the services) the defaultcustomerservice
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}

func (c DefaultCustomerService) GetCustomerById(id int) (domain.Customer, error) {
	return c.repo.GetCustomerById(id)
}
