package repositories

import (
	"cash_register/internal/domain/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	BeginTransaction(ctx context.Context) (*gorm.DB, error)
	CommitTransaction(tx *gorm.DB) error
	RollbackTransaction(tx *gorm.DB) error
	CreateTransaction(tx *gorm.DB, transaction *models.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

// BeginTransaction inicia una transacción
func (r *transactionRepository) BeginTransaction(ctx context.Context) (*gorm.DB, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

// CommitTransaction permite confirmar una transacción
func (r *transactionRepository) CommitTransaction(tx *gorm.DB) error {
	if tx == nil {
		return errors.New("no active transaction")
	}
	return tx.Commit().Error
}

// RollbackTransaction - no inserta nada en caso de error
func (r *transactionRepository) RollbackTransaction(tx *gorm.DB) error {
	if tx == nil {
		return errors.New("no active transaction")
	}
	return tx.Rollback().Error
}

// CreateTransaction inserta una nueva transacción en la base de datos
func (r *transactionRepository) CreateTransaction(tx *gorm.DB, transaction *models.Transaction) error {
	if tx == nil {
		return errors.New("no active transaction")
	}
	return tx.Create(transaction).Error
}
