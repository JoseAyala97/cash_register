package dtos

type PaymentDTO struct {
	TransactionTypeId int                     `json:"transactionTypeId"`
	PaidDetails       []DenominationDetailDTO `json:"paidDetails"`
	Details           []DenominationDetailDTO `json:"details"`
}

type DenominationDetailDTO struct {
	DenominationId int `json:"denominationId"`
	Quantity       int `json:"quantity"`
}
