package dto

type TransactionRequest struct {
	Id int `json:"-"`
	UserEmail string `json:"-"`
	InvoiceNumber string `json:"-"`
	ServiceCode string `json:"service_code" validate:"required"`
	ServiceName string `json:"-"`
	Description string `json:"-"`
	TransactionType string `json:"-"`
	TotalAmount float64 `json:"-"`
	CreatedOn string `json:"-"`
}

type TransactionResponse struct {
	Id int `json:"-"`
	UserEmail string `json:"-"`
	InvoiceNumber string `json:"invoice_number"`
	ServiceCode string `json:"service_code"`
	ServiceName string `json:"service_name"`
	TransactionType string `json:"transaction_type"`
	Description string `json:"-"`
	TotalAmount float64 `json:"total_amount"`
	CreatedOn string `json:"created_on"`
}

type TransactionHistoryResponse struct {
	Id int `json:"-"`
	UserEmail string `json:"-"`
	InvoiceNumber string `json:"invoice_number"`
	ServiceCode string `json:"-"`
	ServiceName string `json:"-"`
	TransactionType string `json:"transaction_type"`
	Description string `json:"description"`
	TotalAmount float64 `json:"total_amount"`
	CreatedOn string `json:"created_on"`	
}