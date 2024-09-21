package dtos

type TransactionDTO struct {
	TransactionTypeId int                    `json:"transactionTypeId" binding:"required"`
	TotalAmount       float64                `json:"totalAmount"`
	PaidAmount        float64                `json:"paidAmount"`
	Details           []TransactionDetailDTO `json:"details"`
}

// DTO para los detalles de la transacción (sin TotalAmount, se calculará automáticamente)
type TransactionDetailDTO struct {
	DenominationId int `json:"denominationId" binding:"required"`
	Quantity       int `json:"quantity" binding:"required"`
}
