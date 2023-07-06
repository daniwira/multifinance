package repository

import (
	"github.com/daniwira/multifinance/internal/domain/loan"
	"gorm.io/gorm"
)

type LoanRepository interface {
	Create(loan *loan.Loan) error
	Update(loan *loan.Loan) error
	Delete(loan *loan.Loan) error
	FindByID(id uint) (*loan.Loan, error)
	FindByCustomerID(customerID uint) (*loan.Loan, error)
	// Add other repository methods as needed
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{
		db: db,
	}
}

func (r *loanRepository) Create(loan *loan.Loan) error {
	return r.db.Create(loan).Error
}

func (r *loanRepository) Update(loanParams *loan.Loan) error {
	return r.db.Model(&loan.Loan{}).Where("id = ?", loanParams.ID).Updates(loanParams).Error
}

func (r *loanRepository) Delete(loan *loan.Loan) error {
	return r.db.Delete(loan).Error
}

func (r *loanRepository) FindByID(id uint) (*loan.Loan, error) {
	var loan loan.Loan
	err := r.db.First(&loan, id).Error
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (r *loanRepository) FindByCustomerID(customerID uint) (*loan.Loan, error) {
	var loan loan.Loan
	err := r.db.Where("customer_id = ?", customerID).First(&loan).Error
	if err != nil {
		return nil, err
	}
	return &loan, nil
}
