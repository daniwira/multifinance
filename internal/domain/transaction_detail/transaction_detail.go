package transactiondetail

type TransactionDetail struct {
	ID                     uint    `gorm:"primaryKey"`
	LimitTenor1            float64 `gorm:"column:limit_tenor1"`
	LimitTenor2            float64 `gorm:"column:limit_tenor2"`
	LimitTenor3            float64 `gorm:"column:limit_tenor3"`
	LimitTenor4            float64 `gorm:"column:limit_tenor4"`
	CustomerNIK            string  `gorm:"column:customer_nik"`
	CustomerFullName       string  `gorm:"column:customer_full_name"`
	CustomerLegalName      string  `gorm:"column:customer_legal_name"`
	CustomerPlaceOfBirth   string  `gorm:"column:customer_place_of_birth"`
	CustomerDateOfBirth    string  `gorm:"column:customer_date_of_birth"`
	CustomerSalary         float64 `gorm:"column:customer_salary"`
	CustomerID             uint    `gorm:"column:customer_id"`
	LoanID                 uint    `gorm:"column:loan_id"`
	LoanInstallment        float64 `gorm:"column:loan_installment"`
	LoanInterestPercentage float64 `gorm:"column:loan_interest_persentage"`
	LoanTotalLoan          float64 `gorm:"column:loan_total_loan"`
	LoanInterest           float64 `gorm:"column:loan_interest"`
	LoanAdminFee           float64 `gorm:"column:loan_admin_fee"`
	LoanOTR                float64 `gorm:"column:loan_otr"`
}

func (TransactionDetail) TableName() string {
	return "transaction_detail"
}
