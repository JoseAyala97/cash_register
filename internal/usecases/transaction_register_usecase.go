package usecases

import (
	"cash_register/internal/adapters/dtos"
	"cash_register/internal/adapters/repositories"
	"cash_register/internal/domain/models"
	"cash_register/internal/interfaces"
	"context"
	"errors"
)

type TransactionRegister struct {
	transactionRepo       repositories.TransactionRepository
	transactionDetailRepo repositories.TransactionDetailRepository
	currentRegisterRepo   repositories.CurrentRegisterRepository
	denominationRepo      interfaces.DenominationRepository
}

// NewTransactionRegister es el constructor del caso de uso de registro de transacciones
func NewTransactionRegister(tr repositories.TransactionRepository, tdr repositories.TransactionDetailRepository, crr repositories.CurrentRegisterRepository, dr interfaces.DenominationRepository) *TransactionRegister {
	return &TransactionRegister{
		transactionRepo:       tr,
		transactionDetailRepo: tdr,
		currentRegisterRepo:   crr,
		denominationRepo:      dr,
	}
}
func (uc *TransactionRegister) RegisterTransaction(ctx context.Context, transaction *models.Transaction, details []dtos.TransactionDetailDTO) error {
	// Iniciar una transacción
	tx, err := uc.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	// Calcular el TotalAmount basado en las denominaciones y la cantidad
	totalAmount := 0.0
	var transactionDetails []*models.TransactionDetail

	for _, detailDTO := range details {
		// Obtener la denominación desde el repositorio
		denomination, err := uc.denominationRepo.GetByID(uint(detailDTO.DenominationId))
		if err != nil {
			uc.transactionRepo.RollbackTransaction(tx)
			return errors.New("denominación no encontrada")
		}

		// Calcular el TotalAmount para este detalle
		amount := denomination.Value * float64(detailDTO.Quantity)
		totalAmount += amount

		// Crear un nuevo detalle de transacción
		transactionDetails = append(transactionDetails, &models.TransactionDetail{
			DenominationId: detailDTO.DenominationId,
			Quantity:       detailDTO.Quantity,
			TotalAmount:    amount, // El monto calculado
		})
	}

	// Asignar el TotalAmount calculado a la transacción principal
	transaction.TotalAmount = totalAmount

	// Registrar la transacción principal
	err = uc.transactionRepo.CreateTransaction(tx, transaction)
	if err != nil {
		uc.transactionRepo.RollbackTransaction(tx)
		return err
	}

	// Ahora que la transacción principal tiene un ID asignado, registrar los detalles
	for _, detail := range transactionDetails {
		detail.TransactionId = transaction.Id // Asignar el ID de la transacción principal a cada detalle
		err = uc.transactionDetailRepo.CreateTransactionDetail(tx, detail)
		if err != nil {
			uc.transactionRepo.RollbackTransaction(tx)
			return err
		}

		// Actualizar el registro actual de caja (current_registers)
		err = uc.currentRegisterRepo.UpdateRegister(tx, detail.DenominationId, detail.Quantity)
		if err != nil {
			uc.transactionRepo.RollbackTransaction(tx)
			return err
		}
	}

	// Confirmar la transacción
	err = uc.transactionRepo.CommitTransaction(tx)
	if err != nil {
		return err
	}

	return nil
}
