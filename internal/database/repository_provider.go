package database

import (
	"context"
	"pismo-back-teste/internal"
)

// RepositoryProvider returns repositories scoped to the given context.
type RepositoryProvider struct{}

func NewRepositoryProvider() *RepositoryProvider {
	return &RepositoryProvider{}
}

func (p *RepositoryProvider) AccountRepository(ctx context.Context) internal.IAccountRepository {
	return NewAccountRepository(MustGetTx(ctx))
}

func (p *RepositoryProvider) TransactionRepository(ctx context.Context) internal.ITransactionRepository {
	return NewTransactionRepository(MustGetTx(ctx))
}
