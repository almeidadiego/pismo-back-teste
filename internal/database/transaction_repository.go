package database

import (
	"context"
	"database/sql"
	"errors"
	"pismo-back-teste/internal/api/dto"
	"time"
)

type TransactionRepository struct {
	tx *sql.Tx
}

func NewTransactionRepository(tx *sql.Tx) *TransactionRepository {
	return &TransactionRepository{tx: tx}
}

// CreateTransaction inserts an account to database
func (r *TransactionRepository) CreateTransaction(
	ctx context.Context, transaction dto.Transaction) error {
	query := `INSERT INTO transaction (account_id, operation_type_id, amount, created_at) values (?,?,?,?)`

	createdAt := time.Now().Format("2006-01-02 15:04:05")
	res, err := r.tx.ExecContext(
		ctx,
		query,
		transaction.AccountID,
		transaction.OperationTypeID,
		transaction.Amount,
		createdAt,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return errors.New("no rows were affected")
	}

	return nil
}
