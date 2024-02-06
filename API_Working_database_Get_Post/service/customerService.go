package service

import "api/domain"

// CustomerService interface defines the methods for managing customers
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(id int) (*domain.Customer, error)
	CreateCustomer(newCustomer domain.Customer) (*domain.Customer, error)
}

// DefaultCustomerService is the implementation of the CustomerService interface
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// GetAllCustomers retrieves all customers
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// NewCustomerService creates a new instance of DefaultCustomerService
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}

// GetCustomerById retrieves a customer by ID
func (c DefaultCustomerService) GetCustomerById(id int) (*domain.Customer, error) {
	return c.repo.GetCustomerById(id)
}

// CreateCustomer creates a new customer
func (s DefaultCustomerService) CreateCustomer(newCustomer domain.Customer) (*domain.Customer, error) {
	createdCustomer, err := s.repo.CreateCustomer(newCustomer)
	if err != nil {
		return nil, err
	}
	return createdCustomer, nil
}
