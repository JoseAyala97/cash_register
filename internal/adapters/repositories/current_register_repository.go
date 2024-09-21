package repositories

import (
	"cash_register/internal/domain/models"
	"errors"

	"gorm.io/gorm"
)

// CurrentRegisterRepository interface para manejar el estado de la caja
type CurrentRegisterRepository interface {
	UpdateRegister(tx *gorm.DB, denominationID int, quantity int) error
}

type currentRegisterRepository struct {
	db *gorm.DB
}

// NewCurrentRegisterRepository constructor para crear el repositorio de CurrentRegister
func NewCurrentRegisterRepository(db *gorm.DB) CurrentRegisterRepository {
	return &currentRegisterRepository{db: db}
}

// UpdateRegister actualiza la cantidad de una denominación específica en la caja
func (r *currentRegisterRepository) UpdateRegister(tx *gorm.DB, denominationID int, quantity int) error {
	if tx == nil {
		return errors.New("no active transaction")
	}

	// Actualiza la cantidad de una denominación específica en la caja (suma o resta monedas/billetes)
	return tx.Model(&models.CurrentRegister{}).
		Where("denomination_id = ?", denominationID).
		Update("quantity", gorm.Expr("quantity + ?", quantity)).
		Error
}
