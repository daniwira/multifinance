package service

import (
	"fmt"

	"github.com/daniwira/multifinance/internal/domain/customer"
	"github.com/daniwira/multifinance/internal/repository"
)

type CustomerService interface {
	GetCustomers() ([]customer.Customer, error)
	GetCustomer(id string) (*customer.Customer, error)
	CreateCustomer(customer customer.Customer) (*customer.Customer, error)
	UpdateCustomer(customer customer.Customer) (*customer.Customer, error)
	DeleteCustomer(id string) error
}

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

func (s *customerService) GetCustomers() ([]customer.Customer, error) {
	return s.customerRepo.GetCustomers()
}

func (s *customerService) GetCustomer(id string) (*customer.Customer, error) {
	return s.customerRepo.GetCustomer(id)
}

func (s *customerService) CreateCustomer(customer customer.Customer) (*customer.Customer, error) {
	return s.customerRepo.CreateCustomer(customer)
}

func (s *customerService) UpdateCustomer(customer customer.Customer) (*customer.Customer, error) {
	customerID := fmt.Sprintf("%d", customer.ID)
	existingCustomer, err := s.customerRepo.GetCustomer(customerID)
	if err != nil {
		return nil, err
	}

	// Perform any necessary validation or business logic before updating
	existingCustomer.FullName = customer.FullName
	existingCustomer.LegalName = customer.LegalName
	// Update other fields as needed

	return s.customerRepo.UpdateCustomer(existingCustomer)
}

func (s *customerService) DeleteCustomer(id string) error {
	existingCustomer, err := s.customerRepo.GetCustomer(id)
	if err != nil {
		return err
	}

	return s.customerRepo.DeleteCustomer(existingCustomer)
}
