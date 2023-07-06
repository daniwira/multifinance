package service

import (
	"errors"
	"fmt"

	"github.com/daniwira/multifinance/internal/domain/loan"
	transactiondetail "github.com/daniwira/multifinance/internal/domain/transaction_detail"
	"github.com/daniwira/multifinance/internal/repository"
)

type LoanService interface {
	CreateLoan(loan *loan.LoanParams) error
	UpdateLoan(id uint, loan *loan.Loan) error
	DeleteLoan(id uint) error
	GetLoanByID(id uint) (*loan.Loan, error)
	PaymentInstallment(params *loan.PaymentInstallment) error
}

type loanService struct {
	loanRepository              repository.LoanRepository
	limitRepository             repository.LimitRepository
	customerRepository          repository.CustomerRepository
	transactionDetailRepository repository.TransactionDetailRepository
}

func NewLoanService(loanRepo repository.LoanRepository, limitRepo repository.LimitRepository, customerRepo repository.CustomerRepository,
	transactionDetRepo repository.TransactionDetailRepository) LoanService {
	return &loanService{
		loanRepository:              loanRepo,
		limitRepository:             limitRepo,
		customerRepository:          customerRepo,
		transactionDetailRepository: transactionDetRepo,
	}
}

func (s *loanService) CreateLoan(params *loan.LoanParams) error {
	var (
		adminFee = 2.00
	)

	limitID := fmt.Sprintf("%d", params.LimitID)
	limit, err := s.limitRepository.GetLimit(limitID)
	if err != nil {
		return err
	}

	tenorLimit, err := limit.GetTenorValue(params.TotalMonth)
	if err != nil {
		return err
	}

	if params.TotalLoan > tenorLimit {
		return errors.New("out of range")
	}

	customerID := fmt.Sprintf("%d", params.CustomerID)
	customer, err := s.customerRepository.GetCustomer(customerID)
	if err != nil {
		return err
	}

	installment := params.TotalLoan + (params.TotalLoan*params.InterestPercentage)/float64(params.TotalMonth)
	interest := params.TotalLoan * params.InterestPercentage

	loanPayload := &loan.Loan{
		TotalLoan:          params.TotalLoan,
		CustomerID:         uint(customer.ID),
		LimitID:            uint(limit.ID),
		InterestPercentage: params.InterestPercentage,
		Installment:        installment,
		Interest:           interest,
		AdminFee:           adminFee,
		OTR:                params.OTR,
	}

	return s.loanRepository.Create(loanPayload)
}

func (s *loanService) PaymentInstallment(params *loan.PaymentInstallment) error {
	customerID := fmt.Sprintf("%d", params.CustomerID)
	customer, err := s.customerRepository.GetCustomer(customerID)
	if err != nil {
		return err
	}

	cID := uint(customer.ID)
	loan, err := s.loanRepository.FindByCustomerID(cID)
	if err != nil {
		return err
	}

	if loan.Installment != params.Amount {
		return errors.New("installment amount not match")
	}

	limitID := fmt.Sprintf("%d", loan.LimitID)
	limit, err := s.limitRepository.GetLimit(limitID)
	if err != nil {
		return err
	}

	transactionDetail := &transactiondetail.TransactionDetail{
		LimitTenor1:            limit.Tenor1,
		LimitTenor2:            limit.Tenor2,
		LimitTenor3:            limit.Tenor3,
		LimitTenor4:            limit.Tenor4,
		CustomerNIK:            customer.Nik,
		CustomerFullName:       customer.FullName,
		CustomerLegalName:      customer.LegalName,
		CustomerPlaceOfBirth:   customer.PlaceOfBirth,
		CustomerDateOfBirth:    customer.DateOfBirth,
		CustomerSalary:         customer.Salary,
		CustomerID:             uint(customer.ID),
		LoanID:                 loan.ID,
		LoanInstallment:        loan.Installment,
		LoanInterestPercentage: loan.InterestPercentage,
		LoanTotalLoan:          loan.TotalLoan,
		LoanInterest:           loan.Interest,
		LoanAdminFee:           loan.AdminFee,
		LoanOTR:                loan.OTR,
	}
	return s.transactionDetailRepository.Create(transactionDetail)
}

func (s *loanService) UpdateLoan(id uint, loan *loan.Loan) error {
	existingLoan, err := s.loanRepository.FindByID(id)
	if err != nil {
		return err
	}
	existingLoan.TotalLoan = loan.TotalLoan
	existingLoan.CustomerID = loan.CustomerID
	existingLoan.LimitID = loan.LimitID
	existingLoan.InterestPercentage = loan.InterestPercentage

	return s.loanRepository.Update(existingLoan)
}

func (s *loanService) DeleteLoan(id uint) error {
	loan, err := s.loanRepository.FindByID(id)
	if err != nil {
		return err
	}

	return s.loanRepository.Delete(loan)
}

func (s *loanService) GetLoanByID(id uint) (*loan.Loan, error) {
	return s.loanRepository.FindByID(id)
}
