package service

import (
	"api/domain"
	"api/dto"
)

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service api/service CustomerService

// CustomerService interface defines the methods for managing customers
type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, error)
	GetCustomerById(id int) (*dto.CustomerResponse, error)
	CreateCustomer(newCustomer domain.Customer) (*dto.CustomerResponse, error)
}

// DefaultCustomerService is the implementation of the CustomerService interface
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers retrieves all customers
func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, error) {
	// retrive all the customers from the repository
	c, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	// Create an empty slice to store the responses
	var responses []dto.CustomerResponse

	//Iterate over each customer returned by the repository
	for _, customer := range c {
		// Convert each customer to a CustomerResponse using the ToDO() Method
		response := customer.ToDto()
		// Append the CustomerResponse to the response slice
		responses = append(responses, response)

	}

	return responses, nil

}

// NewCustomerService creates a new instance of DefaultCustomerService
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}

// GetCustomerById retrieves a customer by ID
func (s DefaultCustomerService) GetCustomerById(id int) (*dto.CustomerResponse, error) {
	c, err := s.repo.GetCustomerById(id)
	if err != nil {
		return nil, err

	}

	c.ToDto()

	response := c.ToDto()
	return &response, nil
}

// CreateCustomer creates a new customer
func (s DefaultCustomerService) CreateCustomer(newCustomer domain.Customer) (*dto.CustomerResponse, error) {
	createdCustomer, err := s.repo.CreateCustomer(newCustomer)
	if err != nil {
		return nil, err
	}

	response := createdCustomer.ToDto()
	return &response, nil
}
