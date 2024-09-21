package repositories

import (
	"cash_register/internal/domain/models"
	"errors"

	"gorm.io/gorm"
)

type TransactionDetailRepository interface {
	CreateTransactionDetail(tx *gorm.DB, detail *models.TransactionDetail) error
}

type transactionDetailRepository struct {
	db *gorm.DB
}

func NewTransactionDetailRepository(db *gorm.DB) TransactionDetailRepository {
	return &transactionDetailRepository{db: db}
}

// Crear un detalle de transacción
func (r *transactionDetailRepository) CreateTransactionDetail(tx *gorm.DB, detail *models.TransactionDetail) error {
	if tx == nil {
		return errors.New("no active transaction")
	}
	return tx.Create(detail).Error
}
