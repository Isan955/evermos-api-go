package usecase

import (
	"errors"
	"evermos-api/internal/entity"
	"evermos-api/internal/repository"
)

type TransactionUsecase interface {
	Create(userID uint, trx *entity.Transaction) error
	GetMyTransactions(userID uint) ([]entity.Transaction, error)
	GetByID(id uint, userID uint) (*entity.Transaction, error)
	Cancel(id uint, userID uint) error
}

type transactionUsecase struct {
	productRepo     repository.ProductRepository
	transactionRepo repository.TransactionRepository
}

func NewTransactionUsecase(p repository.ProductRepository, t repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{productRepo: p, transactionRepo: t}
}

func (u *transactionUsecase) Create(userID uint, trx *entity.Transaction) error {
	trx.UserID = userID
	return u.transactionRepo.Create(trx)
}

func (u *transactionUsecase) GetMyTransactions(userID uint) ([]entity.Transaction, error) {
	return u.transactionRepo.GetMyTransactions(userID)
}

func (u *transactionUsecase) GetByID(id uint, userID uint) (*entity.Transaction, error) {
	trx, err := u.transactionRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if trx.UserID != userID {
		return nil, errors.New("not your transaction")
	}

	return trx, nil
}

func (u *transactionUsecase) Cancel(id uint, userID uint) error {
	trx, err := u.transactionRepo.GetByID(id)
	if err != nil {
		return err
	}

	if trx.UserID != userID {
		return errors.New("access denied")
	}

	trx.Status = "canceled"
	return u.transactionRepo.Update(trx)
}
