package service

import (
	"fmt"
	"nutech-test/internal/dto"
	"nutech-test/internal/entity"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
	GetAllTransactionByEmail(email string) ([]entity.Transaction, error)
	GetTransactionByInvoice(invoice string) (entity.Transaction, error)
}

type ServicesRepository interface {
	GetServiceByCode(serviceCode string) (entity.Service, error)
}

type TransactionServ struct {
	transactionRepository TransactionRepository
	serviceRepos ServicesRepository
}

func NewTransactionService(tr TransactionRepository, sr ServicesRepository) *TransactionServ {
	return &TransactionServ{transactionRepository: tr, serviceRepos: sr}
}

func (ts *TransactionServ) CreateTransactione(req dto.TransactionRequest, email string) (dto.TransactionResponse, error) {
	//get service info 
	services, err := ts.serviceRepos.GetServiceByCode(req.ServiceCode)
	if err != nil {
		return dto.TransactionResponse{}, fmt.Errorf("get services info %s", err)
	}

	//using uuid more safely because its unique and we dont have to handle the race condition/concurrency
	invoiceNumber := uuid.NewString()

	request := entity.Transaction{
		UserEmail: email,
		ServiceCode: req.ServiceCode,
		ServiceName: services.ServiceName,
		InvoiceNumber: invoiceNumber,
		TransactionType: "PAYMENT",
		Description: services.ServiceName,
		TotalAmount: services.ServiceTariff,
	}
	//create
	if err := ts.transactionRepository.Create(&request); err != nil {
		return dto.TransactionResponse{}, fmt.Errorf("create transaction %s", err)
	}

	transaction, err := ts.GetTransactionByInvoice(invoiceNumber)
	if err != nil {
		return dto.TransactionResponse{}, fmt.Errorf("get transaction by invoice %s", err)
	}

	resp := dto.TransactionResponse(transaction)

	return resp, nil
}

func (ts *TransactionServ) GetAllTransactionByEmail(email string) ([]dto.TransactionHistoryResponse, error) {
	transactions, err := ts.transactionRepository.GetAllTransactionByEmail(email)
	if err != nil {
		return []dto.TransactionHistoryResponse{}, fmt.Errorf("get all transaction %s", err)
	}

	var resp []dto.TransactionHistoryResponse
	for _, transaction := range transactions {
		resp = append(resp, dto.TransactionHistoryResponse(transaction))
	}

	return resp, nil
}

func (ts *TransactionServ) GetTransactionByInvoice(invoice string) (dto.TransactionResponse, error) {
	transaction, err := ts.transactionRepository.GetTransactionByInvoice(invoice)
	if err != nil {
		return dto.TransactionResponse{}, fmt.Errorf("get transaction by invoice %s", err)
	}

	resp := dto.TransactionResponse(transaction)

	return resp, nil
}