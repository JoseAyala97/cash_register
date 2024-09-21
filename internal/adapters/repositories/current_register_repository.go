package repositories

import (
	"cash_register/internal/adapters/views"
	"cash_register/internal/domain/models"
	"errors"

	"gorm.io/gorm"
)

type CurrentRegisterRepository interface {
	UpdateRegister(tx *gorm.DB, denominationID int, quantity int) error
	GetCurrentRegister() ([]views.CurrentRegisterView, float64, error)
}

type currentRegisterRepository struct {
	db *gorm.DB
}

func NewCurrentRegisterRepository(db *gorm.DB) CurrentRegisterRepository {
	return &currentRegisterRepository{db: db}
}

// Actualizar o insertar un registro en la tabla current_registers
func (r *currentRegisterRepository) UpdateRegister(tx *gorm.DB, denominationID int, quantity int) error {
	if tx == nil {
		return errors.New("no active transaction")
	}

	// Verificar si ya existe un registro para la denominación
	var currentRegister models.CurrentRegister
	err := tx.Where("denomination_id = ?", denominationID).First(&currentRegister).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err == gorm.ErrRecordNotFound {
		// Si no existe, insertar un nuevo registro en CurrentRegister
		newRegister := models.CurrentRegister{
			DenominationId: denominationID,
			Quantity:       quantity,
		}
		return tx.Create(&newRegister).Error
	}

	// Si existe, actualizar la cantidad (sumar la nueva cantidad)
	return tx.Model(&currentRegister).Update("quantity", gorm.Expr("quantity + ?", quantity)).Error
}

// Obtener el estado actual de la caja
func (r *currentRegisterRepository) GetCurrentRegister() ([]views.CurrentRegisterView, float64, error) {
	var currentRegister []models.CurrentRegister
	err := r.db.Preload("Denomination.MoneyType").Find(&currentRegister).Error
	if err != nil {
		return nil, 0, err
	}

	// Mapear los resultados a CurrentRegisterView y calcular el total acumulado
	var currentRegisterViews []views.CurrentRegisterView
	var totalGeneral float64

	for _, register := range currentRegister {
		// Calcular el total: cantidad * valor de la denominación
		total := calculateTotal(register.Quantity, register.Denomination.Value)

		// Sumar al total general
		totalGeneral += total

		// Agregar a la lista de vistas con el total calculado para cada denominación
		currentRegisterViews = append(currentRegisterViews, views.CurrentRegisterView{
			Id:       register.Id,
			Quantity: register.Quantity,
			DenominationView: views.DenominationView{
				Id:    register.Denomination.Id,
				Value: register.Denomination.Value,
				MoneyType: views.MoneyTypeView{
					Id:    register.Denomination.MoneyType.Id,
					Value: register.Denomination.MoneyType.Value,
				},
			},
			Total: total, // Total calculado para esta denominación
		})
	}

	return currentRegisterViews, totalGeneral, nil
}

// Método privado para calcular el total de una denominación
func calculateTotal(quantity int, denominationValue float64) float64 {
	return float64(quantity) * denominationValue
}
