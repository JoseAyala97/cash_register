package interfaces

import "cash_register/internal/domain/models"

type MoneyTypeRepository interface {
	GetAll() ([]models.MoneyType, error)
	GetByID(id uint) (*models.MoneyType, error)
	Create(moneyType models.MoneyType) error
	Update(id uint, moneyType models.MoneyType) error
	Delete(id uint) error
}
