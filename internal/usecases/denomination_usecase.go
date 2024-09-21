package usecases

import (
	"cash_register/internal/adapters/views"
	"cash_register/internal/domain/models"
	"cash_register/internal/interfaces"
)

type DenominationUsecase struct {
	repo interfaces.DenominationRepository // Cambiar a DenominationRepository
}

// Constructor del caso de uso para Denomination
func NewDenominationUsecase(repo interfaces.DenominationRepository) *DenominationUsecase {
	return &DenominationUsecase{repo: repo}
}

func (uc *DenominationUsecase) CreateDenomination(denomination models.Denomination) error {
	return uc.repo.Create(denomination)
}

func (uc *DenominationUsecase) GetAllDenominations() ([]views.DenominationView, error) {
	// Obtenemos todas las denominaciones del repositorio
	denominations, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Creamos un slice para almacenar las vistas
	var denominationViews []views.DenominationView

	// Convertimos cada modelo Denomination a DenominationView
	for _, denomination := range denominations {
		denominationView := views.DenominationView{
			Id:    denomination.Id,
			Value: denomination.Value,
			MoneyType: views.MoneyTypeView{
				Id:    denomination.MoneyType.Id,
				Value: denomination.MoneyType.Value,
			},
		}
		denominationViews = append(denominationViews, denominationView)
	}

	// Retornamos el slice de vistas
	return denominationViews, nil
}

func (uc *DenominationUsecase) GetDenominationByID(id uint) (*models.Denomination, error) {
	return uc.repo.GetByID(id)
}

func (uc *DenominationUsecase) UpdateDenomination(id uint, updatedDenomination models.Denomination) error {
	return uc.repo.Update(id, updatedDenomination)
}

func (uc *DenominationUsecase) DeleteDenomination(id uint) error {
	return uc.repo.Delete(id)
}
