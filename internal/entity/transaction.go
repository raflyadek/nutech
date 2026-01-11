package entity

type Transaction struct {
	Id int
	UserEmail string
	InvoiceNumber string
	ServiceCode string
	ServiceName string
	TransactionType string
	Description string
	TotalAmount float64
	CreatedOn string
}
