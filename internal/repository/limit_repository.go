package repository

import (
	"github.com/daniwira/multifinance/internal/domain/limit"
	"gorm.io/gorm"
)

type LimitRepository interface {
	GetLimits() ([]limit.Limit, error)
	GetLimit(id string) (*limit.Limit, error)
	CreateLimit(limit limit.Limit) (*limit.Limit, error)
	UpdateLimit(limit *limit.Limit) (*limit.Limit, error)
	DeleteLimit(limit *limit.Limit) error
	FindByCustomerID(customerID uint) (*limit.Limit, error)
}

type limitRepository struct {
	db *gorm.DB
}

func NewLimitRepository(db *gorm.DB) LimitRepository {
	return &limitRepository{
		db: db,
	}
}

func (r *limitRepository) GetLimits() ([]limit.Limit, error) {
	var limits []limit.Limit
	result := r.db.Find(&limits)
	if result.Error != nil {
		return nil, result.Error
	}
	return limits, nil
}

func (r *limitRepository) GetLimit(id string) (*limit.Limit, error) {
	var limit limit.Limit
	result := r.db.First(&limit, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &limit, nil
}

func (r *limitRepository) CreateLimit(limit limit.Limit) (*limit.Limit, error) {
	result := r.db.Create(&limit)
	if result.Error != nil {
		return nil, result.Error
	}
	return &limit, nil
}

func (r *limitRepository) UpdateLimit(limit *limit.Limit) (*limit.Limit, error) {
	result := r.db.Save(&limit)
	if result.Error != nil {
		return nil, result.Error
	}
	return limit, nil
}

func (r *limitRepository) DeleteLimit(limit *limit.Limit) error {
	result := r.db.Delete(&limit)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *limitRepository) FindByCustomerID(customerID uint) (*limit.Limit, error) {
	var limit limit.Limit
	err := r.db.Where("customer_id = ?", customerID).First(&limit).Error
	if err != nil {
		return nil, err
	}
	return &limit, nil
}
