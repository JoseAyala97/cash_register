package views

type DenominationView struct {
	Id        int           `json:"id"`
	Value     float64       `json:"value"`
	MoneyType MoneyTypeView `json:"moneyType"`
}
