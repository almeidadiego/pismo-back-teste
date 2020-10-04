package internal

import (
	"context"
	"pismo-back-teste/internal/api/dto"
)

type IRepositoryProvider interface {
	AccountRepository(ctx context.Context) IAccountRepository
	TransactionRepository(ctx context.Context) ITransactionRepository
}

type IAccountRepository interface {
	GetAccount(ctx context.Context, id int) (dto.Account, error)
	CreateAccount(ctx context.Context, account dto.Account) error
}

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction dto.Transaction) error
}
