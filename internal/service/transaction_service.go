package service

import (
	"errors"
	"fmt"
	"nutech-test/internal/dto"
	"nutech-test/internal/entity"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
	GetAllTransactionByEmail(email string, limit int, offset int) ([]entity.Transaction, error)
	GetTransactionByInvoice(invoice string) (entity.Transaction, error)
}

type UserRepos interface {
	GetBalanceByEmail(email string) (entity.Saldo, error)
	UpdateBalanceByEmail(saldo *entity.Saldo) (entity.Saldo, error)
}
type ServicesRepository interface {
	GetServiceByCode(serviceCode string) (entity.Service, error)
}

type TransactionServ struct {
	transactionRepository TransactionRepository
	serviceRepos ServicesRepository
	userRepos UserRepos
}

func NewTransactionService(tr TransactionRepository, sr ServicesRepository, ur UserRepos) *TransactionServ {
	return &TransactionServ{transactionRepository: tr, serviceRepos: sr, userRepos: ur}
}

func (ts *TransactionServ) CreateTransaction(req dto.TransactionRequest, email string) (dto.TransactionResponse, error) {
	//check balance first
	balance, err := ts.userRepos.GetBalanceByEmail(email)
	if err != nil {
		return dto.TransactionResponse{}, fmt.Errorf("get balance by email %s", err)
	}

	//get service info 
	services, err := ts.serviceRepos.GetServiceByCode(req.ServiceCode)
	if err != nil {
		return dto.TransactionResponse{}, fmt.Errorf("get services info %s", err)
	}

	//validate if service exists
	if req.ServiceCode != services.ServiceCode {
		return dto.TransactionResponse{}, errors.New("Service atau Layanan tidak ditemukan")
	}

	//check if balance enough to buy a service
	if balance.Balance < services.ServiceTariff {
		return dto.TransactionResponse{}, errors.New("balance is not enough")
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

	//if success then balance user - total amount 
	balanceAfter := balance.Balance - services.ServiceTariff
	balanceRequest := entity.Saldo{
		UserEmail: email,
		Balance: balanceAfter,
	}
	//update balance
	_, errr := ts.userRepos.UpdateBalanceByEmail(&balanceRequest)
	if errr != nil {
		return dto.TransactionResponse{}, fmt.Errorf("update balance %s", err)
	}

	resp := dto.TransactionResponse(transaction)

	return resp, nil
}

func (ts *TransactionServ) GetAllTransactionByEmail(
	email string,
	limit int,
	offset int,
) ([]dto.TransactionHistoryResponse, error) {

	transactions, err := ts.transactionRepository.
		GetAllTransactionByEmail(email, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("get all transaction: %w", err)
	}

	resp := make([]dto.TransactionHistoryResponse, 0, len(transactions))
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