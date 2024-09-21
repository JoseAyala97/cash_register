package interfaces

import "cash_register/internal/domain/models"

type DenominationRepository interface {
	GetAll() ([]models.Denomination, error)
	GetByID(id uint) (*models.Denomination, error)
	Create(denomination models.Denomination) error
	Update(id uint, denomination models.Denomination) error
	Delete(id uint) error
}
