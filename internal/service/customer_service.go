package service

import (
	"fmt"

	domaincustomer "github.com/daniwira/multifinance/internal/domain/customer"
	"github.com/daniwira/multifinance/internal/repository"
	"github.com/microcosm-cc/bluemonday"
)

type CustomerService interface {
	GetCustomers() ([]domaincustomer.Customer, error)
	GetCustomer(id string) (*domaincustomer.Customer, error)
	CreateCustomer(customer domaincustomer.Customer) (*domaincustomer.Customer, error)
	UpdateCustomer(customer domaincustomer.Customer) (*domaincustomer.Customer, error)
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

func (s *customerService) GetCustomers() ([]domaincustomer.Customer, error) {
	return s.customerRepo.GetCustomers()
}

func (s *customerService) GetCustomer(id string) (*domaincustomer.Customer, error) {
	return s.customerRepo.GetCustomer(id)
}

func (s *customerService) CreateCustomer(customer domaincustomer.Customer) (*domaincustomer.Customer, error) {
	return s.customerRepo.CreateCustomer(customer)
}

func (s *customerService) UpdateCustomer(customer domaincustomer.Customer) (*domaincustomer.Customer, error) {
	customerID := fmt.Sprintf("%d", customer.ID)
	existingCustomer, err := s.customerRepo.GetCustomer(customerID)
	if err != nil {
		return nil, err
	}

	sanitizeCustomer(&customer)
	// Perform any necessary validation or business logic before updating
	existingCustomer.FullName = customer.FullName
	existingCustomer.LegalName = customer.LegalName

	return s.customerRepo.UpdateCustomer(existingCustomer)
}

func (s *customerService) DeleteCustomer(id string) error {
	existingCustomer, err := s.customerRepo.GetCustomer(id)
	if err != nil {
		return err
	}

	return s.customerRepo.DeleteCustomer(existingCustomer)
}

func sanitizeCustomer(customer *domaincustomer.Customer) {
	p := bluemonday.UGCPolicy()

	customer.FullName = p.Sanitize(customer.FullName)
	customer.LegalName = p.Sanitize(customer.LegalName)

}
