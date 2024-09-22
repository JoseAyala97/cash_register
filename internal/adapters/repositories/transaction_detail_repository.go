package repositories

import (
	"cash_register/internal/domain/models"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type TransactionDetailRepository interface {
	CreateTransactionDetail(tx *gorm.DB, detail *models.TransactionDetail) error
	GetTransactionLogs(ctx context.Context, startDateTime, endDateTime *time.Time) ([]models.TransactionDetail, error)
}

type transactionDetailRepository struct {
	db *gorm.DB
}

func NewTransactionDetailRepository(db *gorm.DB) *transactionDetailRepository {
	return &transactionDetailRepository{db: db}
}

// Crear un detalle de transacción
func (r *transactionDetailRepository) CreateTransactionDetail(tx *gorm.DB, detail *models.TransactionDetail) error {
	if tx == nil {
		return errors.New("no active transaction")
	}
	return tx.Create(detail).Error
}

// Obtener los logs de transacciones con filtros de fecha
func (r *transactionDetailRepository) GetTransactionLogs(ctx context.Context, startDateTime, endDateTime *time.Time) ([]models.TransactionDetail, error) {
	var transactionDetails []models.TransactionDetail

	// Crear la consulta básica
	query := r.db.WithContext(ctx).Model(&models.TransactionDetail{}).
		Preload("Transaction.TransactionType").
		Preload("Denomination")

	// Aplicar filtro de fecha de inicio, si existe
	if startDateTime != nil {
		query = query.Where("created_at >= ?", *startDateTime)
	}

	// Aplicar filtro de fecha de fin, si existe
	if endDateTime != nil {
		query = query.Where("created_at <= ?", *endDateTime)
	}

	// Ejecutar la consulta y retornar los resultados
	err := query.Find(&transactionDetails).Error
	if err != nil {
		return nil, err
	}

	return transactionDetails, nil
}
