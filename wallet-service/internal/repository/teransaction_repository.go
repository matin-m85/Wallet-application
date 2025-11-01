package repository

import (
	"wallet-service/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(txn *models.Transaction) error {
	return r.db.Create(txn).Error
}

func (r *TransactionRepository) ListByWalletID(walletID string) ([]models.Transaction, error) {
	var list []models.Transaction
	if err := r.db.Where("wallet_id = ?", walletID).Order("created_at desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
