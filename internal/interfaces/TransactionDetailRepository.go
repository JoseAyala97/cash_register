package interfaces

import (
	"cash_register/internal/domain/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type TransactionDetailRepository interface {
	CreateTransactionDetail(tx *gorm.DB, detail *models.TransactionDetail) error
	GetTransactionLogs(ctx context.Context, startDateTime, endDateTime *time.Time) ([]models.TransactionDetail, error)
}
