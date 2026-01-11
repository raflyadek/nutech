package dto

type Service struct {
	Id int `json:"-"`
	ServiceCode string `json:"service_code"`
	ServiceName string `json:"service_name"`
	ServiceIcon string `json:"service_icon"`
	ServiceTariff float64 `json:"service_tariff"`
}