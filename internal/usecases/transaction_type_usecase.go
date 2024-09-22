package usecases

import (
	"cash_register/internal/domain/models"
	"cash_register/internal/interfaces"
)

type TransactionTypeUsecase struct {
	repo interfaces.Repository[models.TransactionType]
}

// NewTransactionTypeUsecase es el constructor para el caso de uso de TransactionType
func NewTransactionTypeUsecase(repo interfaces.Repository[models.TransactionType]) *TransactionTypeUsecase {
	return &TransactionTypeUsecase{repo: repo}
}

// Crear un nuevo tipo de transacción
func (uc *TransactionTypeUsecase) CreateTransactionType(transactionType models.TransactionType) error {
	return uc.repo.Create(transactionType)
}

// Obtener todos los tipos de transacciones
func (uc *TransactionTypeUsecase) GetAllTransactionTypes() ([]models.TransactionType, error) {
	return uc.repo.GetAll()
}

// Obtener un tipo de transacción por ID
func (uc *TransactionTypeUsecase) GetTransactionTypeByID(id uint) (*models.TransactionType, error) {
	return uc.repo.GetByID(id)
}

// Actualizar un tipo de transacción
func (uc *TransactionTypeUsecase) UpdateTransactionType(id uint, transactionType models.TransactionType) error {
	return uc.repo.Update(id, transactionType)
}

// Eliminar un tipo de transacción
func (uc *TransactionTypeUsecase) DeleteTransactionType(id uint) error {
	return uc.repo.Delete(id)
}
