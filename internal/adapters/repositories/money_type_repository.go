package repositories

import (
	"cash_register/internal/domain/models"
	"cash_register/internal/interfaces"

	"gorm.io/gorm"
)

// Implementación del MoneyTypeRepository
type MoneyTypeRepositoryImpl struct {
	db *gorm.DB
}

// Asegúrate de que implementa la interfaz
var _ interfaces.MoneyTypeRepository = &MoneyTypeRepositoryImpl{}

// Constructor para crear una instancia del repositorio
func NewMoneyTypeRepository(db *gorm.DB) *MoneyTypeRepositoryImpl {
	return &MoneyTypeRepositoryImpl{db: db}
}

// Crear un nuevo tipo de moneda
func (r *MoneyTypeRepositoryImpl) Create(moneyType models.MoneyType) error {
	return r.db.Create(&moneyType).Error
}

// Obtener todos los tipos de moneda
func (r *MoneyTypeRepositoryImpl) GetAll() ([]models.MoneyType, error) {
	var moneyTypes []models.MoneyType
	if err := r.db.Find(&moneyTypes).Error; err != nil {
		return nil, err
	}
	return moneyTypes, nil
}

// Obtener un tipo de moneda por ID
func (r *MoneyTypeRepositoryImpl) GetByID(id uint) (*models.MoneyType, error) {
	var moneyType models.MoneyType
	if err := r.db.First(&moneyType, id).Error; err != nil {
		return nil, err
	}
	return &moneyType, nil
}

// Actualizar un tipo de moneda
func (r *MoneyTypeRepositoryImpl) Update(id uint, updatedMoneyType models.MoneyType) error {
	return r.db.Model(&models.MoneyType{}).Where("id = ?", id).Updates(updatedMoneyType).Error
}

// Eliminar un tipo de moneda
func (r *MoneyTypeRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.MoneyType{}, id).Error
}
