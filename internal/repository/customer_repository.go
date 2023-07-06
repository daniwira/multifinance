package repository

import (
	"github.com/daniwira/multifinance/internal/domain/customer"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomers() ([]customer.Customer, error)
	GetCustomer(id string) (*customer.Customer, error)
	CreateCustomer(customer customer.Customer) (*customer.Customer, error)
	UpdateCustomer(customer *customer.Customer) (*customer.Customer, error)
	DeleteCustomer(customer *customer.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) GetCustomers() ([]customer.Customer, error) {
	var customers []customer.Customer
	result := r.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (r *customerRepository) GetCustomer(id string) (*customer.Customer, error) {
	var customer customer.Customer
	result := r.db.Debug().First(&customer, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r *customerRepository) CreateCustomer(customer customer.Customer) (*customer.Customer, error) {
	result := r.db.Create(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r *customerRepository) UpdateCustomer(customer *customer.Customer) (*customer.Customer, error) {
	result := r.db.Debug().Save(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}

func (r *customerRepository) DeleteCustomer(customer *customer.Customer) error {
	result := r.db.Delete(&customer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
