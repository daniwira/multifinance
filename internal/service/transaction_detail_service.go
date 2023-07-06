package service

import (
	transactiondetail "github.com/daniwira/multifinance/internal/domain/transaction_detail"
	"github.com/daniwira/multifinance/internal/repository"
)

type TransactionDetailService interface {
	CreateTransactionDetail(transaction *transactiondetail.TransactionDetail) error
	DeleteTransactionDetail(id uint) error
	GetTransactionDetailByID(id uint) (*transactiondetail.TransactionDetail, error)
}

type transactionDetailService struct {
	transactionDetailRepository repository.TransactionDetailRepository
}

func NewTransactionDetailService(transactionDetailRepo repository.TransactionDetailRepository) TransactionDetailService {
	return &transactionDetailService{
		transactionDetailRepository: transactionDetailRepo,
	}
}

func (s *transactionDetailService) CreateTransactionDetail(transaction *transactiondetail.TransactionDetail) error {
	return s.transactionDetailRepository.Create(transaction)
}

func (s *transactionDetailService) DeleteTransactionDetail(id uint) error {
	transaction, err := s.transactionDetailRepository.FindByID(id)
	if err != nil {
		return err
	}

	return s.transactionDetailRepository.Delete(transaction)
}

func (s *transactionDetailService) GetTransactionDetailByID(id uint) (*transactiondetail.TransactionDetail, error) {
	return s.transactionDetailRepository.FindByID(id)
}
