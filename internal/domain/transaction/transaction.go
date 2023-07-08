package domaintransaction

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID           uint64         `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	CustomerID   uint           `gorm:"column:customer_id" json:"customer_id"`
	ContractNo   string         `gorm:"column:contract_no" json:"contract_no"`
	OTR          float64        `gorm:"column:otr" json:"otr"`
	AdminFee     float64        `gorm:"column:admin_fee" json:"admin_fee"`
	Installments int            `gorm:"column:installments" json:"installments"`
	Interest     float64        `gorm:"column:interest" json:"interest"`
	AssetName    string         `gorm:"column:asset_name" json:"asset_name"`
}
