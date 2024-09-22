package usecases

import (
	"cash_register/internal/adapters/repositories"
	"cash_register/internal/adapters/views"
	"context"
	"time"
)

type TransactionDetailUsecase struct {
	transactionDetailRepo repositories.TransactionDetailRepository
}

// Constructor para el caso de uso
func NewTransactionDetailUsecase(tdr repositories.TransactionDetailRepository) *TransactionDetailUsecase {
	return &TransactionDetailUsecase{
		transactionDetailRepo: tdr,
	}
}

// Obtener los logs de transacciones con los filtros de fecha
func (uc *TransactionDetailUsecase) GetTransactionLogs(ctx context.Context, startDateTime, endDateTime *time.Time) ([]views.TransactionDetailView, error) {
	transactionDetails, err := uc.transactionDetailRepo.GetTransactionLogs(ctx, startDateTime, endDateTime)
	if err != nil {
		return nil, err
	}

	var transactionDetailViews []views.TransactionDetailView
	for _, detail := range transactionDetails {
		transactionDetailView := views.TransactionDetailView{
			Id: detail.Id,
			DenominationView: views.DenominationView{
				Id:    detail.Denomination.Id,
				Value: detail.Denomination.Value,
				MoneyType: views.MoneyTypeView{
					Id:    detail.Denomination.MoneyType.Id,
					Value: detail.Denomination.MoneyType.Value,
				},
			},
			Quantity:    detail.Quantity,
			TotalAmount: detail.TotalAmount,
			TransactionView: views.TransactionView{
				Id:          detail.Transaction.Id,
				TotalAmount: detail.Transaction.TotalAmount,
				PaidAmount:  detail.Transaction.PaidAmount,
				TransactionTypeView: views.TransactionTypeView{
					Id:    detail.Transaction.TransactionType.Id,
					Value: detail.Transaction.TransactionType.Value,
				},
			},
		}
		transactionDetailViews = append(transactionDetailViews, transactionDetailView)
	}

	totalInEvent, totalOutEvent := uc.calculateTotalInAndOut(transactionDetailViews)

	for i := range transactionDetailViews {
		transactionDetailViews[i].TotalInEvent = totalInEvent
		transactionDetailViews[i].TotalOutEvent = totalOutEvent
	}

	return transactionDetailViews, nil
}

func (uc *TransactionDetailUsecase) calculateTotalInAndOut(details []views.TransactionDetailView) (float64, float64) {
	var totalInEvent, totalOutEvent float64

	for _, detail := range details {
		if detail.TransactionView.TransactionTypeView.Id == 1 {
			totalInEvent += detail.TransactionView.TotalAmount
		} else if detail.TransactionView.TransactionTypeView.Id == 4 {
			totalOutEvent += detail.TransactionView.TotalAmount
		}
	}

	return totalInEvent, totalOutEvent
}
