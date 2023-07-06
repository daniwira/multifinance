package repository

import (
	"github.com/daniwira/multifinance/internal/domain/transaction"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactions() ([]transaction.Transaction, error)
	GetTransaction(id string) (*transaction.Transaction, error)
	CreateTransaction(transaction transaction.Transaction) (*transaction.Transaction, error)
	UpdateTransaction(transaction *transaction.Transaction) (*transaction.Transaction, error)
	DeleteTransaction(transaction *transaction.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) GetTransactions() ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction
	result := r.db.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (r *transactionRepository) GetTransaction(id string) (*transaction.Transaction, error) {
	var transaction transaction.Transaction
	result := r.db.First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

func (r *transactionRepository) CreateTransaction(transaction transaction.Transaction) (*transaction.Transaction, error) {
	result := r.db.Create(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

func (r *transactionRepository) UpdateTransaction(transaction *transaction.Transaction) (*transaction.Transaction, error) {
	result := r.db.Save(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return transaction, nil
}

func (r *transactionRepository) DeleteTransaction(transaction *transaction.Transaction) error {
	result := r.db.Delete(&transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
