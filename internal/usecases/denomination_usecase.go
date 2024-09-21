package usecases

import (
	"cash_register/internal/domain/models"
	"cash_register/internal/interfaces"
)

type DenominationUsecase struct {
	repo interfaces.DenominationRepository
}

// Constructor para crear el DenominationUsecase inyectando el repositorio
func NewDenominationUsecase(repo interfaces.DenominationRepository) *DenominationUsecase {
	return &DenominationUsecase{repo: repo}
}

// Obtener todas las denominaciones
func (uc *DenominationUsecase) GetAllDenominations() ([]models.Denomination, error) {
	return uc.repo.GetAll()
}

// Crear una nueva denominación
func (uc *DenominationUsecase) CreateDenomination(denomination models.Denomination) error {
	return uc.repo.Create(denomination)
}

// Obtener una denominación por ID
func (uc *DenominationUsecase) GetDenominationByID(id uint) (*models.Denomination, error) {
	return uc.repo.GetByID(id)
}

// Actualizar una denominación por ID
func (uc *DenominationUsecase) UpdateDenomination(id uint, updatedDenomination models.Denomination) error {
	return uc.repo.Update(id, updatedDenomination)
}

// Eliminar una denominación por ID
func (uc *DenominationUsecase) DeleteDenomination(id uint) error {
	return uc.repo.Delete(id)
}
