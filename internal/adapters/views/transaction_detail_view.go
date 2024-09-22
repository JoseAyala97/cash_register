package views

type TransactionDetailView struct {
	Id               int              `json:"id"`
	DenominationView DenominationView `json:"denomination"`
	Quantity         int              `json:"quantity"`
	TotalAmount      float64          `json:"totalAmount"`
	TransactionView  TransactionView  `json:"transaction"`
	TotalInEvent     float64          `json:"totalInEvent"`
	TotalOutEvent    float64          `json:"totalOutEvent"`
}
