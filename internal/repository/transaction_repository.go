package repository

import (
	domaintransaction "github.com/daniwira/multifinance/internal/domain/transaction"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactions() ([]domaintransaction.Transaction, error)
	GetTransaction(id string) (*domaintransaction.Transaction, error)
	CreateTransaction(transaction domaintransaction.Transaction) (*domaintransaction.Transaction, error)
	UpdateTransaction(transaction *domaintransaction.Transaction) (*domaintransaction.Transaction, error)
	DeleteTransaction(transaction *domaintransaction.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) GetTransactions() ([]domaintransaction.Transaction, error) {
	var transactions []domaintransaction.Transaction
	result := r.db.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (r *transactionRepository) GetTransaction(id string) (*domaintransaction.Transaction, error) {
	var transaction domaintransaction.Transaction
	result := r.db.First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

func (r *transactionRepository) CreateTransaction(transaction domaintransaction.Transaction) (*domaintransaction.Transaction, error) {
	result := r.db.Create(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

func (r *transactionRepository) UpdateTransaction(transaction *domaintransaction.Transaction) (*domaintransaction.Transaction, error) {
	result := r.db.Save(&transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return transaction, nil
}

func (r *transactionRepository) DeleteTransaction(transaction *domaintransaction.Transaction) error {
	result := r.db.Delete(&transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
