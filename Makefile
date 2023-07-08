test:
	@go test -v ./...
mock:
	/Users/apple/go/bin/mockgen --source=internal/repository/customer_repository.go -destination=shared/mock/repository/customer_repository.go --package mock
	/Users/apple/go/bin/mockgen --source=internal/repository/limit_repository.go -destination=shared/mock/repository/limit_repository.go --package mock
	/Users/apple/go/bin/mockgen --source=internal/repository/loan.go -destination=shared/mock/repository/loan.go --package mock
	/Users/apple/go/bin/mockgen --source=internal/repository/transaction_detail.go -destination=shared/mock/repository/transaction_detail.go --package mock
	/Users/apple/go/bin/mockgen --source=internal/repository/transaction_repository.go -destination=shared/mock/repository/transaction_repository.go --package mock