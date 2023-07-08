package service

import (
	"errors"
	"testing"

	domaintransaction "github.com/daniwira/multifinance/internal/domain/transaction"
	mock "github.com/daniwira/multifinance/shared/mock/repository"
	"go.uber.org/mock/gomock"

	"github.com/smartystreets/goconvey/convey"
)

func TestGetTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	convey.Convey("Given a GetTransactions", t, func() {
		transactionRepo := mock.NewMockTransactionRepository(ctrl)
		service := NewTransactionService(transactionRepo)

		convey.Convey("When GetTransactions is called", func() {
			respTransactions := []domaintransaction.Transaction{
				{ID: 1},
				{ID: 2},
			}
			transactionRepo.EXPECT().GetTransactions().Return(respTransactions, nil)

			resp, err := service.GetTransactions()
			convey.Convey("Then it should return the expected transactions and no error", func() {
				convey.So(err, convey.ShouldBeNil)
				convey.So(resp, convey.ShouldNotBeEmpty)
			})
		})
	})

	convey.Convey("Given a GetTransactions with error", t, func() {
		transactionRepo := mock.NewMockTransactionRepository(ctrl)
		service := NewTransactionService(transactionRepo)

		convey.Convey("When GetTransactions is called", func() {
			respTransactions := []domaintransaction.Transaction{}
			transactionRepo.EXPECT().GetTransactions().Return(respTransactions, errors.New("error"))

			resp, err := service.GetTransactions()
			convey.Convey("Then it should return the expected transactions and with error", func() {
				convey.So(err, convey.ShouldNotBeNil)
				convey.So(resp, convey.ShouldBeEmpty)
			})
		})
	})

}
