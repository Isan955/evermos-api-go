package repository

import (
	"evermos-api/config"
	"evermos-api/internal/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
	GetByID(id uint) (*entity.Transaction, error)
	GetMyTransactions(userID uint) ([]entity.Transaction, error)
	Update(transaction *entity.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{
		db: config.DB,
	}
}

func (r *transactionRepository) Create(transaction *entity.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetByID(id uint) (*entity.Transaction, error) {
	var trx entity.Transaction
	err := r.db.Preload("TransactionItems").
		Where("id = ?", id).
		First(&trx).Error
	return &trx, err
}

func (r *transactionRepository) GetMyTransactions(userID uint) ([]entity.Transaction, error) {
	var trxs []entity.Transaction
	err := r.db.Preload("TransactionItems").
		Where("user_id = ?", userID).
		Find(&trxs).Error
	return trxs, err
}

func (r *transactionRepository) Update(transaction *entity.Transaction) error {
	return r.db.Save(transaction).Error
}
