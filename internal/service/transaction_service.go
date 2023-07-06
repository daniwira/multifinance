package service

import (
	"fmt"

	"github.com/daniwira/multifinance/internal/domain/transaction"
	"github.com/daniwira/multifinance/internal/repository"
)

type TransactionService interface {
	GetTransactions() ([]transaction.Transaction, error)
	GetTransaction(id string) (*transaction.Transaction, error)
	CreateTransaction(transaction transaction.Transaction) (*transaction.Transaction, error)
	UpdateTransaction(transaction transaction.Transaction) (*transaction.Transaction, error)
	DeleteTransaction(id string) error
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}

func (s *transactionService) GetTransactions() ([]transaction.Transaction, error) {
	return s.transactionRepo.GetTransactions()
}

func (s *transactionService) GetTransaction(id string) (*transaction.Transaction, error) {
	return s.transactionRepo.GetTransaction(id)
}

func (s *transactionService) CreateTransaction(transaction transaction.Transaction) (*transaction.Transaction, error) {
	return s.transactionRepo.CreateTransaction(transaction)
}

func (s *transactionService) UpdateTransaction(transaction transaction.Transaction) (*transaction.Transaction, error) {
	transactionID := fmt.Sprintf("%d", transaction.ID)
	existingTransaction, err := s.transactionRepo.GetTransaction(transactionID)
	if err != nil {
		return nil, err
	}

	// Perform any necessary validation or business logic before updating
	existingTransaction.ContractNo = transaction.ContractNo
	existingTransaction.OTR = transaction.OTR
	// Update other fields as needed

	return s.transactionRepo.UpdateTransaction(existingTransaction)
}

func (s *transactionService) DeleteTransaction(id string) error {
	existingTransaction, err := s.transactionRepo.GetTransaction(id)
	if err != nil {
		return err
	}

	return s.transactionRepo.DeleteTransaction(existingTransaction)
}
