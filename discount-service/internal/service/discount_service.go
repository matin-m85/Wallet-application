package service

import (
	"context"
	"discount-servise/internal/models"
	"discount-servise/internal/repository"

	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DiscountService struct {
	repo *repository.DiscountRepository
}

func NewDiscountService(db *gorm.DB) *DiscountService {
	return &DiscountService{repo: repository.NewDiscountRepository(db)}
}

// Create discount
func (s *DiscountService) Create(ctx context.Context, code, desc string, percent int, remaining int) (*models.Discount, error) {
	d := &models.Discount{
		ID:          uuid.New(),
		Code:        code,
		Description: desc,
		Percent:     percent,
		Remaining:   remaining,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := s.repo.Create(d); err != nil {
		return nil, err
	}
	return d, nil
}

// Get by code
func (s *DiscountService) GetByCode(ctx context.Context, code string) (*models.Discount, error) {
	return s.repo.GetByCode(code)
}

// Decrement remaining (simple redeem)
func (s *DiscountService) Redeem(ctx context.Context, code string) (*models.Discount, error) {
	d, err := s.repo.GetByCode(code)
	if err != nil {
		return nil, err
	}
	if !d.IsActive || d.Remaining <= 0 {
		return nil, fmt.Errorf("invalid or exhausted discount")
	}
	d.Remaining -= 1
	d.UpdatedAt = time.Now()
	if d.Remaining == 0 {
		d.IsActive = false
	}
	if err := s.repo.Update(d); err != nil {
		return nil, err
	}
	return d, nil
}

func (s *DiscountService) ListAll() ([]models.Discount, error) {
	return s.repo.ListAll()
}
