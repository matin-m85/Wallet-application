package repository

import (
	"discount-servise/internal/models"

	"gorm.io/gorm"
)

type DiscountRepository struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) *DiscountRepository {
	return &DiscountRepository{db: db}
}

func (r *DiscountRepository) Create(d *models.Discount) error {
	return r.db.Create(d).Error
}

func (r *DiscountRepository) GetByCode(code string) (*models.Discount, error) {
	var d models.Discount
	if err := r.db.Where("code = ?", code).First(&d).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *DiscountRepository) Update(d *models.Discount) error {
	return r.db.Save(d).Error
}

func (r *DiscountRepository) ListAll() ([]models.Discount, error) {
	var list []models.Discount
	if err := r.db.Order("created_at desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
