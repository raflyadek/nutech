package repository

import (
	"context"
	"database/sql"
	"nutech-test/internal/entity"
)

type TransactionRepo struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (tr *TransactionRepo) Create(transaction *entity.Transaction) error {
	_, err := tr.db.ExecContext(context.Background(), `
	INSERT INTO transaction (user_email, invoice_number, service_code, service_name, description, transaction_type, total_amount)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`, transaction.UserEmail, transaction.InvoiceNumber, transaction.ServiceCode, transaction.ServiceName, transaction.Description, transaction.TransactionType, transaction.TotalAmount)

	if err != nil {
		return err
	}

	return nil
}

func (tr *TransactionRepo) GetAllTransactionByEmail(email string) ([]entity.Transaction, error) {
	rows, err := tr.db.QueryContext(context.Background(), `
	SELECT invoice_number, transaction_type, description, total_amount, created_on FROM transaction`)
	if err != nil {
		return []entity.Transaction{}, err
	}
	defer rows.Close()

	var transactions []entity.Transaction

	for rows.Next() {
		var t entity.Transaction
		if err := rows.Scan(&t.InvoiceNumber, &t.TransactionType, &t.Description, &t.TotalAmount, &t.CreatedOn); err != nil {
			return []entity.Transaction{}, err
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (tr *TransactionRepo) GetTransactionByInvoice(invoice string) (entity.Transaction, error) {
	var transaction entity.Transaction
	err := tr.db.QueryRowContext(context.Background(), `
	SELECT invoice_number, service_code, service_name, transaction_type, total_amount, created_on FROM transaction WHERE invoice_number = $1`, invoice,
	).Scan(&transaction.InvoiceNumber, &transaction.ServiceCode, &transaction.ServiceName, &transaction.TransactionType, &transaction.TotalAmount, &transaction.CreatedOn)

	if err != nil {
		return entity.Transaction{}, err
	}

	return transaction, nil
}