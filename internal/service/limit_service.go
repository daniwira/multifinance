package service

import (
	"fmt"

	domainlimit "github.com/daniwira/multifinance/internal/domain/limit"
	"github.com/daniwira/multifinance/internal/repository"
)

type LimitService interface {
	GetLimits() ([]domainlimit.Limit, error)
	GetLimit(id string) (*domainlimit.Limit, error)
	CreateLimit(limit domainlimit.Limit) (*domainlimit.Limit, error)
	UpdateLimit(limit domainlimit.Limit) (*domainlimit.Limit, error)
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

func (s *limitService) GetLimits() ([]domainlimit.Limit, error) {
	return s.limitRepo.GetLimits()
}

func (s *limitService) GetLimit(id string) (*domainlimit.Limit, error) {
	return s.limitRepo.GetLimit(id)
}

func (s *limitService) CreateLimit(limit domainlimit.Limit) (*domainlimit.Limit, error) {
	return s.limitRepo.CreateLimit(limit)
}

func (s *limitService) UpdateLimit(limit domainlimit.Limit) (*domainlimit.Limit, error) {
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
