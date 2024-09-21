package usecases

import (
	"cash_register/internal/adapters/repositories"
	"cash_register/internal/adapters/views"
)

type CurrentRegisterUsecase struct {
	currentRegisterRepo repositories.CurrentRegisterRepository
}

// NewCurrentRegisterUsecase es el constructor para el caso de uso de CurrentRegister
func NewCurrentRegisterUsecase(crr repositories.CurrentRegisterRepository) *CurrentRegisterUsecase {
	return &CurrentRegisterUsecase{
		currentRegisterRepo: crr,
	}
}

// Obtener el estado actual de la caja y el total general
func (uc *CurrentRegisterUsecase) GetCurrentRegister() ([]views.CurrentRegisterView, float64, error) {
	return uc.currentRegisterRepo.GetCurrentRegister()
}
