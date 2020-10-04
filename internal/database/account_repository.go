package database

import (
	"context"
	"database/sql"
	"errors"
	"pismo-back-teste/internal/api/dto"
)

type AccountRepository struct {
	tx *sql.Tx
}

func NewAccountRepository(tx *sql.Tx) *AccountRepository {
	return &AccountRepository{tx: tx}
}

func (r *AccountRepository) GetAccount(ctx context.Context, id int) (dto.Account, error) {
	account := dto.Account{}

	query := `SELECT id, document FROM account WHERE id = ?`

	err := r.tx.QueryRowContext(ctx, query, id).Scan(&account.ID, &account.Document)

	switch {
	case err == nil || err == sql.ErrNoRows:
		return account, nil // 200 or 404
	default:
		return account, err // 500
	}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, account dto.Account) error {
	query := `INSERT INTO account (document) values (?)`

	res, err := r.tx.ExecContext(ctx, query, account.Document)
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
