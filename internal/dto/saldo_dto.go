package dto

type TopUpSaldoRequest struct {
	Balance float64 `json:"top_up_amount" validate:"required,numeric"`
}

type SaldoResponse struct {
	Id int `json:"-"`
	UserEmail string `json:"-"`
	Balance float64 `json:"balance"`
}