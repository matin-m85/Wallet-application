package repository

import (
	"wallet-service/internal/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) GetByPhone(phone string) (*models.Wallet, error) {
	var w models.Wallet
	if err := r.db.Where("phone = ?", phone).First(&w).Error; err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *WalletRepository) Create(w *models.Wallet) error {
	return r.db.Create(w).Error
}

func (r *WalletRepository) Save(w *models.Wallet) error {
	return r.db.Save(w).Error
}

func (r *WalletRepository) ListAll() ([]models.Wallet, error) {
	var list []models.Wallet
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
