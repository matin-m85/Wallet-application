package service

import (
	"context"
	"errors"
	"time"
	"wallet-service/internal/models"
	"wallet-service/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletService struct {
	db         *gorm.DB
	walletRepo *repository.WalletRepository
	txnRepo    *repository.TransactionRepository
}

func NewWalletService(db *gorm.DB) *WalletService {
	return &WalletService{
		db:         db,
		walletRepo: repository.NewWalletRepository(db),
		txnRepo:    repository.NewTransactionRepository(db),
	}
}

func (s *WalletService) GetOrCreateWalletByPhone(ctx context.Context, phone string) (*models.Wallet, error) {
	// Try get
	w, err := s.walletRepo.GetByPhone(phone)
	if err == nil {
		return w, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	// create
	w = &models.Wallet{
		ID:      uuid.New(),
		Phone:   phone,
		Balance: 0,
	}
	if err := s.walletRepo.Create(w); err != nil {
		return nil, err
	}
	return w, nil
}

func (s *WalletService) GetBalance(ctx context.Context, phone string) (int64, error) {
	w, err := s.GetOrCreateWalletByPhone(ctx, phone)
	if err != nil {
		return 0, err
	}
	return w.Balance, nil
}

func (s *WalletService) ApplyTopUp(ctx context.Context, phone string, amount int64, reference string) (*models.Transaction, error) {
	var createdTxn *models.Transaction
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// lock wallet row
		var w models.Wallet
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("phone = ?", phone).First(&w).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// create
				w = models.Wallet{
					ID:      uuid.New(),
					Phone:   phone,
					Balance: 0,
				}
				if err := tx.Create(&w).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}

		// update balance
		w.Balance += amount
		if err := tx.Save(&w).Error; err != nil {
			return err
		}

		// transaction record
		txn := &models.Transaction{
			ID:        uuid.New(),
			WalletID:  w.ID,
			Amount:    amount,
			Reference: reference,
			CreatedAt: time.Now(),
		}
		if err := tx.Create(txn).Error; err != nil {
			return err
		}
		createdTxn = txn
		return nil
	})
	if err != nil {
		return nil, err
	}
	return createdTxn, nil
}

func (s *WalletService) ListTransactions(ctx context.Context, phone string) ([]models.Transaction, error) {
	w, err := s.GetOrCreateWalletByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	return s.txnRepo.ListByWalletID(w.ID.String())
}
func (s *WalletService) ListAllWallets() ([]models.Wallet, error) {
	var list []models.Wallet
	if err := s.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
