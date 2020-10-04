package database

import (
	"context"
	"database/sql"
)

type key int

const (
	txkey key = 1
)

func WithTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txkey, tx)
}

func MustGetTx(ctx context.Context) *sql.Tx {
	tx, ok := getTx(ctx)
	if !ok {
		panic("context has no sql.Tx value")
	}
	return tx
}

func getTx(ctx context.Context) (*sql.Tx, bool) {
	queryer, ok := ctx.Value(txkey).(*sql.Tx)
	return queryer, ok
}
