package repository

import (
	domainlimit "github.com/daniwira/multifinance/internal/domain/limit"
	"gorm.io/gorm"
)

type LimitRepository interface {
	GetLimits() ([]domainlimit.Limit, error)
	GetLimit(id string) (*domainlimit.Limit, error)
	CreateLimit(limit domainlimit.Limit) (*domainlimit.Limit, error)
	UpdateLimit(limit *domainlimit.Limit) (*domainlimit.Limit, error)
	DeleteLimit(limit *domainlimit.Limit) error
	FindByCustomerID(customerID uint) (*domainlimit.Limit, error)
}

type limitRepository struct {
	db *gorm.DB
}

func NewLimitRepository(db *gorm.DB) LimitRepository {
	return &limitRepository{
		db: db,
	}
}

func (r *limitRepository) GetLimits() ([]domainlimit.Limit, error) {
	var limits []domainlimit.Limit
	result := r.db.Find(&limits)
	if result.Error != nil {
		return nil, result.Error
	}
	return limits, nil
}

func (r *limitRepository) GetLimit(id string) (*domainlimit.Limit, error) {
	var limit domainlimit.Limit
	result := r.db.First(&limit, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &limit, nil
}

func (r *limitRepository) CreateLimit(limit domainlimit.Limit) (*domainlimit.Limit, error) {
	result := r.db.Create(&limit)
	if result.Error != nil {
		return nil, result.Error
	}
	return &limit, nil
}

func (r *limitRepository) UpdateLimit(limit *domainlimit.Limit) (*domainlimit.Limit, error) {
	result := r.db.Save(&limit)
	if result.Error != nil {
		return nil, result.Error
	}
	return limit, nil
}

func (r *limitRepository) DeleteLimit(limit *domainlimit.Limit) error {
	result := r.db.Delete(&limit)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *limitRepository) FindByCustomerID(customerID uint) (*domainlimit.Limit, error) {
	var limit domainlimit.Limit
	err := r.db.Where("customer_id = ?", customerID).First(&limit).Error
	if err != nil {
		return nil, err
	}
	return &limit, nil
}
