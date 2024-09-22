package views

type CurrentRegisterView struct {
	Id               int              `json:"id"`
	DenominationView DenominationView `json:"denomination"`
	Quantity         int              `json:"quantity"`
	Total            float64          `json:"total"`
}
