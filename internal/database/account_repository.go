package database

import (
	"context"
	"database/sql"
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
