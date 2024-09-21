package dtos

type TransactionDTO struct {
	TransactionTypeId int                    `json:"transactionTypeId" binding:"required"`
	Details           []TransactionDetailDTO `json:"details" binding:"required"`
}

// DTO para los detalles de la transacción (sin TotalAmount, se calculará automáticamente)
type TransactionDetailDTO struct {
	DenominationId int `json:"denominationId" binding:"required"`
	Quantity       int `json:"quantity" binding:"required"`
}
