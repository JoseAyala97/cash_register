package repositories

import (
	"cash_register/internal/domain/models"

	"gorm.io/gorm"
)

type TransactionTypeRepository struct {
	GenericRepository[models.TransactionType]
}

// NewTransactionTypeRepository es un constructor para crear el repositorio de TransactionType
func NewTransactionTypeRepository(db *gorm.DB) *TransactionTypeRepository {
	return &TransactionTypeRepository{
		GenericRepository: *NewGenericRepository[models.TransactionType](db),
	}
}
