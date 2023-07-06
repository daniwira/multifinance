package service

import (
	"fmt"

	"github.com/daniwira/multifinance/internal/domain/limit"
	"github.com/daniwira/multifinance/internal/repository"
)

type LimitService interface {
	GetLimits() ([]limit.Limit, error)
	GetLimit(id string) (*limit.Limit, error)
	CreateLimit(limit limit.Limit) (*limit.Limit, error)
	UpdateLimit(limit limit.Limit) (*limit.Limit, error)
	DeleteLimit(id string) error
}

type limitService struct {
	limitRepo repository.LimitRepository
}

func NewLimitService(limitRepo repository.LimitRepository) LimitService {
	return &limitService{
		limitRepo: limitRepo,
	}
}

func (s *limitService) GetLimits() ([]limit.Limit, error) {
	return s.limitRepo.GetLimits()
}

func (s *limitService) GetLimit(id string) (*limit.Limit, error) {
	return s.limitRepo.GetLimit(id)
}

func (s *limitService) CreateLimit(limit limit.Limit) (*limit.Limit, error) {
	return s.limitRepo.CreateLimit(limit)
}

func (s *limitService) UpdateLimit(limit limit.Limit) (*limit.Limit, error) {
	limitID := fmt.Sprintf("%d", limit.ID)
	existingLimit, err := s.limitRepo.GetLimit(limitID)
	if err != nil {
		return nil, err
	}

	// Perform any necessary validation or business logic before updating
	existingLimit.Name = limit.Name
	existingLimit.Tenor1 = limit.Tenor1
	existingLimit.Tenor2 = limit.Tenor2
	existingLimit.Tenor3 = limit.Tenor3
	existingLimit.Tenor4 = limit.Tenor4
	// Update other fields as needed

	return s.limitRepo.UpdateLimit(existingLimit)
}

func (s *limitService) DeleteLimit(id string) error {
	existingLimit, err := s.limitRepo.GetLimit(id)
	if err != nil {
		return err
	}

	return s.limitRepo.DeleteLimit(existingLimit)
}
