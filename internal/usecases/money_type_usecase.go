package usecases

import (
	"cash_register/internal/domain/models"
	"cash_register/internal/interfaces"
)

type MoneyTypeUsecase struct {
	repo interfaces.Repository[models.MoneyType]
}

// Constructor del caso de uso para MoneyType
func NewMoneyTypeUsecase(repo interfaces.Repository[models.MoneyType]) *MoneyTypeUsecase {
	return &MoneyTypeUsecase{repo: repo}
}

func (uc *MoneyTypeUsecase) CreateMoneyType(moneyType models.MoneyType) error {
	return uc.repo.Create(moneyType)
}

func (uc *MoneyTypeUsecase) GetAllMoneyTypes() ([]models.MoneyType, error) {
	return uc.repo.GetAll()
}

func (uc *MoneyTypeUsecase) GetMoneyTypeByID(id uint) (*models.MoneyType, error) {
	return uc.repo.GetByID(id)
}

func (uc *MoneyTypeUsecase) UpdateMoneyType(id uint, updatedMoneyType models.MoneyType) error {
	return uc.repo.Update(id, updatedMoneyType)
}

func (uc *MoneyTypeUsecase) DeleteMoneyType(id uint) error {
	return uc.repo.Delete(id)
}
