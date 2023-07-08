package repository

import (
	domaincustomer "github.com/daniwira/multifinance/internal/domain/customer"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomers() ([]domaincustomer.Customer, error)
	GetCustomer(id string) (*domaincustomer.Customer, error)
	CreateCustomer(customer domaincustomer.Customer) (*domaincustomer.Customer, error)
	UpdateCustomer(customer *domaincustomer.Customer) (*domaincustomer.Customer, error)
	DeleteCustomer(customer *domaincustomer.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) GetCustomers() ([]domaincustomer.Customer, error) {
	var customers []domaincustomer.Customer
	result := r.db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (r *customerRepository) GetCustomer(id string) (*domaincustomer.Customer, error) {
	var customer domaincustomer.Customer
	result := r.db.Debug().First(&customer, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r *customerRepository) CreateCustomer(customer domaincustomer.Customer) (*domaincustomer.Customer, error) {
	result := r.db.Create(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r *customerRepository) UpdateCustomer(customer *domaincustomer.Customer) (*domaincustomer.Customer, error) {
	result := r.db.Debug().Save(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}

func (r *customerRepository) DeleteCustomer(customer *domaincustomer.Customer) error {
	result := r.db.Delete(&customer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
