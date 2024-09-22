package usecases

import (
	"cash_register/internal/adapters/dtos"
	"cash_register/internal/adapters/repositories"
	"cash_register/internal/domain/models"
	"cash_register/internal/interfaces"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
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

func (uc *TransactionRegister) RegisterTransaction(ctx context.Context, transactionDTO dtos.TransactionDTO) error {
	// Iniciar la transacción en la base de datos
	tx, err := uc.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	// Calcular el monto total basado en las denominaciones y la cantidad
	totalAmount := 0.0
	var transactionDetails []*models.TransactionDetail

	for _, detailDTO := range transactionDTO.Details {
		// Obtener la denominación desde el repositorio
		denomination, err := uc.denominationRepo.GetByID(uint(detailDTO.DenominationId))
		if err != nil {
			uc.transactionRepo.RollbackTransaction(tx)
			return errors.New("denominación no encontrada")
		}

		amount := denomination.Value * float64(detailDTO.Quantity)
		totalAmount += amount

		transactionDetails = append(transactionDetails, &models.TransactionDetail{
			DenominationId: detailDTO.DenominationId,
			Quantity:       detailDTO.Quantity,
			TotalAmount:    amount,
		})
	}

	// Crear la transacción principal
	transaction := models.Transaction{
		TransactionTypeId: transactionDTO.TransactionTypeId,
		TotalAmount:       totalAmount,
	}

	// Evaluar el tipo de transacción basado en el número, excluir pagos
	switch transactionDTO.TransactionTypeId {
	case 1: // "Ingreso Inicial"
		err = uc.handleInitialDeposit(tx, transaction, transactionDetails)
	case 4: // "Retiro Final"
		err = uc.handleWithdrawal(tx)
	default:
		err = errors.New("tipo de transacción no soportado")
	}

	if err != nil {
		uc.transactionRepo.RollbackTransaction(tx)
		return err
	}

	// Confirmar la transacción
	return uc.transactionRepo.CommitTransaction(tx)
}

func (uc *TransactionRegister) handleInitialDeposit(tx *gorm.DB, transaction models.Transaction, details []*models.TransactionDetail) error {
	// Registrar la transacción principal de ingreso inicial
	err := uc.transactionRepo.CreateTransaction(tx, &transaction)
	if err != nil {
		return err
	}

	// Registrar los detalles del ingreso
	for _, detail := range details {
		detail.TransactionId = transaction.Id
		err = uc.transactionDetailRepo.CreateTransactionDetail(tx, detail)
		if err != nil {
			return err
		}

		// Actualizar el registro actual de caja (CurrentRegister)
		err = uc.currentRegisterRepo.UpdateRegister(tx, detail.DenominationId, detail.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}

// Manejo de pagos
func (uc *TransactionRegister) handlePayment(tx *gorm.DB, transaction models.Transaction, paidAmount float64, details []*models.TransactionDetail) error {
	if paidAmount < transaction.TotalAmount {
		return errors.New("el monto pagado es insuficiente")
	}

	// Registrar la transacción principal de pago
	transaction.PaidAmount = paidAmount
	err := uc.transactionRepo.CreateTransaction(tx, &transaction)
	if err != nil {
		return err
	}

	// Registrar los detalles del pago
	for _, detail := range details {
		detail.TransactionId = transaction.Id
		err = uc.transactionDetailRepo.CreateTransactionDetail(tx, detail)
		if err != nil {
			return err
		}
	}

	// Calcular el cambio
	cambio := paidAmount - transaction.TotalAmount
	if cambio > 0 {
		err = uc.handleChange(tx, transaction, cambio)
		if err != nil {
			return err
		}
	}

	return nil
}

// Manejo del cambio
func (uc *TransactionRegister) handleChange(tx *gorm.DB, transaction models.Transaction, change float64) error {
	// Obtener las denominaciones disponibles en la caja registradora, ordenadas de mayor a menor valor
	var currentRegister []models.CurrentRegister
	if err := tx.Preload("Denomination").Find(&currentRegister).Error; err != nil {
		return fmt.Errorf("error al obtener las denominaciones: %v", err)
	}

	// Mapa para almacenar las denominaciones que vamos a devolver
	changeMap := make(map[int]int)

	for _, entry := range currentRegister {
		denomValue := int(entry.Denomination.Value)
		if change == 0 {
			break
		}

		// Determinar cuántas denominaciones de este tipo se pueden usar para el cambio
		count := int(change) / denomValue
		if count > 0 {
			if entry.Quantity >= count {
				changeMap[entry.Denomination.Id] = count
				change -= float64(count * denomValue)
				entry.Quantity -= count
			} else {
				changeMap[entry.Denomination.Id] = entry.Quantity
				change -= float64(entry.Quantity * denomValue)
				entry.Quantity = 0
			}

			// Actualizar la cantidad en la caja registradora
			if err := tx.Save(&entry).Error; err != nil {
				return fmt.Errorf("error al actualizar la caja registradora: %v", err)
			}
		}
	}

	if change > 0 {
		return fmt.Errorf("no hay suficiente cambio disponible en la caja")
	}

	// Registrar los detalles de la transacción para el cambio
	for denomId, qty := range changeMap {
		detail := models.TransactionDetail{
			DenominationId: denomId,
			Quantity:       qty,
			TotalAmount:    float64(qty) * float64(currentRegister[denomId].Denomination.Value),
			TransactionId:  transaction.Id,
		}

		if err := tx.Create(&detail).Error; err != nil {
			return fmt.Errorf("error al registrar los detalles de la transacción: %v", err)
		}
	}

	return nil
}

// Manejo del retiro de la caja
func (uc *TransactionRegister) handleWithdrawal(tx *gorm.DB) error {
	// Obtener el estado actual de la caja
	currentRegister, _, err := uc.currentRegisterRepo.GetCurrentRegister()
	if err != nil {
		return err
	}

	// Iterar sobre todos los registros actuales de la caja
	for _, register := range currentRegister {
		// Marcar el registro como eliminado (soft delete)
		err = tx.Model(&register).Update("deleted_at", gorm.DeletedAt{Time: time.Now(), Valid: true}).Error
		if err != nil {
			return fmt.Errorf("error al marcar el registro de caja como eliminado: %v", err)
		}
	}

	// No se necesita realizar ninguna otra acción o registrar transacciones adicionales.
	return nil
}

func (uc *TransactionRegister) MakePayment(ctx context.Context, paymentDTO dtos.PaymentDTO) error {
	// Iniciar la transacción en la base de datos
	tx, err := uc.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	// Calcular el monto total pagado basado en las denominaciones proporcionadas en `paidDetails`
	totalPaid := 0.0
	for _, paidDetail := range paymentDTO.PaidDetails {
		denomination, err := uc.denominationRepo.GetByID(uint(paidDetail.DenominationId))
		if err != nil {
			uc.transactionRepo.RollbackTransaction(tx)
			return errors.New("denominación no encontrada en el pago")
		}
		totalPaid += denomination.Value * float64(paidDetail.Quantity)
	}

	// Calcular el monto total de la transacción basado en los detalles de la transacción
	totalAmount := 0.0
	var transactionDetails []*models.TransactionDetail
	for _, detailDTO := range paymentDTO.Details {
		denomination, err := uc.denominationRepo.GetByID(uint(detailDTO.DenominationId))
		if err != nil {
			uc.transactionRepo.RollbackTransaction(tx)
			return errors.New("denominación no encontrada en la transacción")
		}

		amount := denomination.Value * float64(detailDTO.Quantity)
		totalAmount += amount

		transactionDetails = append(transactionDetails, &models.TransactionDetail{
			DenominationId: detailDTO.DenominationId,
			Quantity:       detailDTO.Quantity,
			TotalAmount:    amount,
		})
	}

	// Crear la transacción principal (sin detalles ni pago aún)
	transaction := models.Transaction{
		TransactionTypeId: paymentDTO.TransactionTypeId,
		TotalAmount:       totalAmount,
	}

	// Utilizar handlePayment para registrar la transacción, los detalles y manejar el cambio
	err = uc.handlePayment(tx, transaction, totalPaid, transactionDetails)
	if err != nil {
		uc.transactionRepo.RollbackTransaction(tx)
		return err
	}

	// Confirmar la transacción si todo fue exitoso
	return uc.transactionRepo.CommitTransaction(tx)
}
