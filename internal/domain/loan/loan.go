package loan

type Loan struct {
	ID                 uint    `gorm:"primaryKey"`
	TotalLoan          float64 `gorm:"column:total_loan"`
	CustomerID         uint    `gorm:"column:customer_id"`
	LimitID            uint    `gorm:"column:limit_id"`
	InterestPercentage float64 `gorm:"column:interest_persentage"`
	Installment        float64 `gorm:"column:installment"`
	Interest           float64 `gorm:"column:interest"`
	AdminFee           float64 `gorm:"column:admin_fee"`
	OTR                float64 `gorm:"column:otr"`
}

func (Loan) TableName() string {
	return "loan"
}

type LoanParams struct {
	CustomerID         uint    `json:"customer_id"`
	LimitID            uint    `json:"limit_id"`
	TotalMonth         int     `json:"total_month"`
	TotalLoan          float64 `json:"total_loan"`
	OTR                float64 `json:"otr"`
	InterestPercentage float64 `json:"interest_percentage"`
	AssetName          string  `json:"asset_name"`
}

type PaymentInstallment struct {
	CustomerID uint    `json:"customer_id"`
	Amount     float64 `json:"amount"`
}
