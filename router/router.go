package router

import (
	"github.com/daniwira/multifinance/config"
	"github.com/daniwira/multifinance/config/database"
	handlerCustomer "github.com/daniwira/multifinance/internal/domain/handler/customer"
	handlerLimit "github.com/daniwira/multifinance/internal/domain/handler/limit"
	handlerloan "github.com/daniwira/multifinance/internal/domain/handler/loan"
	handlerTransaction "github.com/daniwira/multifinance/internal/domain/handler/transaction"
	"github.com/daniwira/multifinance/internal/repository"
	"github.com/daniwira/multifinance/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (*gin.Engine, error) {
	r := gin.Default()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// Initialize database connection
	db, err := database.InitializeDatabase(cfg.GetDBConnectionString())
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	customerRepo := repository.NewCustomerRepository(db)
	limitRepo := repository.NewLimitRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	loanRepo := repository.NewLoanRepository(db)
	transactionDetailRepo := repository.NewTransactionDetailRepository(db)

	// Initialize services
	customerService := service.NewCustomerService(customerRepo)
	limitService := service.NewLimitService(limitRepo)
	transactionService := service.NewTransactionService(transactionRepo)
	loanSerice := service.NewLoanService(loanRepo, limitRepo, customerRepo, transactionDetailRepo)

	// Initialize handlers
	customerHandler := handlerCustomer.NewCustomerHandler(customerService)
	limitHandler := handlerLimit.NewLimitHandler(limitService)
	transactionHandler := handlerTransaction.NewTransactionHandler(transactionService)
	loanHandler := handlerloan.NewLoanHandler(loanSerice)

	// Define routes
	v1 := r.Group("/api/v1")
	{
		// Customers
		v1.GET("/customers", customerHandler.GetCustomers)
		v1.GET("/customers/:id", customerHandler.GetCustomer)
		v1.POST("/customers", customerHandler.CreateCustomer)
		v1.PUT("/customers/:id", customerHandler.UpdateCustomer)
		v1.DELETE("/customers/:id", customerHandler.DeleteCustomer)

		// Limits
		v1.GET("/limits", limitHandler.GetLimits)
		v1.GET("/limits/:id", limitHandler.GetLimit)
		v1.POST("/limits", limitHandler.CreateLimit)
		v1.PUT("/limits/:id", limitHandler.UpdateLimit)
		v1.DELETE("/limits/:id", limitHandler.DeleteLimit)

		// Transactions
		v1.GET("/transactions", transactionHandler.GetTransactions)
		v1.GET("/transactions/:id", transactionHandler.GetTransaction)
		v1.POST("/transactions", transactionHandler.CreateTransaction)
		v1.PUT("/transactions/:id", transactionHandler.UpdateTransaction)
		v1.DELETE("/transactions/:id", transactionHandler.DeleteTransaction)

		v1.POST("/loan", loanHandler.CreateLoan)
		v1.POST("/loan/payment", loanHandler.PaymentInstallment)
	}

	return r, nil
}

/*
- customer data
- loan
	> tenor type
	> total_month
	> total_loan
	> otr
	> interest_percentage
	> asset_name

	>> jumlah cicilan
	> installment amount= total_loan + (total_loan * interest_percentage )/ total_month

	> otr =  otr
	> interest = total_loan * interest_percentage/100
	> admin fee = 2%
	>
*/
