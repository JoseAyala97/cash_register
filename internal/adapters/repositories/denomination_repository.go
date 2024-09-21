package repositories

import (
	"cash_register/internal/domain/models"
	"cash_register/internal/interfaces"

	"gorm.io/gorm"
)

// Implementación del DenominationRepository
type DenominationRepositoryImpl struct {
	db *gorm.DB
}

// implementacion de la interfaz
var _ interfaces.DenominationRepository = &DenominationRepositoryImpl{}

// Constructor para crear una instancia de DenominationRepositoryImpl
func NewDenominationRepository(db *gorm.DB) *DenominationRepositoryImpl {
	return &DenominationRepositoryImpl{db: db}
}

// Obtener todas las denominaciones
func (r *DenominationRepositoryImpl) GetAll() ([]models.Denomination, error) {
	var denominations []models.Denomination
	err := r.
		db.
		Preload("MoneyType").
		Find(&denominations).
		Error
	if err != nil {
		return nil, err
	}
	return denominations, nil
}

// Crear una nueva denominación
func (r *DenominationRepositoryImpl) Create(denomination models.Denomination) error {
	return r.db.Create(&denomination).Error
}

// Obtener una denominación por ID
func (r *DenominationRepositoryImpl) GetByID(id uint) (*models.Denomination, error) {
	var denomination models.Denomination
	if err := r.db.First(&denomination, id).Error; err != nil {
		return nil, err
	}
	return &denomination, nil
}

// Actualizar una denominación por ID
func (r *DenominationRepositoryImpl) Update(id uint, updatedDenomination models.Denomination) error {
	return r.db.Model(&models.Denomination{}).Where("id = ?", id).Updates(updatedDenomination).Error
}

// Eliminar una denominación por ID
func (r *DenominationRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Denomination{}, id).Error
}
