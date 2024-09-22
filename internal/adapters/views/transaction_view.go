package views

type TransactionView struct {
	Id                  int                 `json:"id"`
	TransactionTypeView TransactionTypeView `json:"transactionType"`
	TotalAmount         float64             `json:"totalAmount"`
	PaidAmount          float64             `json:"paidAmount"`
}

type TransactionTypeView struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}
