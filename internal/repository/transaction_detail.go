package repository

import (
	transactiondetail "github.com/daniwira/multifinance/internal/domain/transaction_detail"
	"gorm.io/gorm"
)

type TransactionDetailRepository interface {
	Create(transaction *transactiondetail.TransactionDetail) error
	Update(transaction *transactiondetail.TransactionDetail) error
	Delete(transaction *transactiondetail.TransactionDetail) error
	FindByID(id uint) (*transactiondetail.TransactionDetail, error)
}

type transactionDetailRepository struct {
	db *gorm.DB
}

func NewTransactionDetailRepository(db *gorm.DB) TransactionDetailRepository {
	return &transactionDetailRepository{
		db: db,
	}
}

func (r *transactionDetailRepository) Create(transaction *transactiondetail.TransactionDetail) error {
	return r.db.Create(transaction).Error
}

func (r *transactionDetailRepository) Update(transaction *transactiondetail.TransactionDetail) error {
	return r.db.Model(&transactiondetail.TransactionDetail{}).Where("id = ?", transaction.ID).Updates(transaction).Error
}

func (r *transactionDetailRepository) Delete(transaction *transactiondetail.TransactionDetail) error {
	return r.db.Delete(transaction).Error
}

func (r *transactionDetailRepository) FindByID(id uint) (*transactiondetail.TransactionDetail, error) {
	var transaction transactiondetail.TransactionDetail
	err := r.db.First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
