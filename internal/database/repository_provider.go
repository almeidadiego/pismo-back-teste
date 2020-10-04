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

// AccountRepository returns a new AccountRepository scoped to the context.
func (p *RepositoryProvider) AccountRepository(ctx context.Context) internal.IAccountRepository {
	return NewAccountRepository(MustGetTx(ctx))
}

// TransactionRepository returns a new TransactionRepository scoped to the context.
func (p *RepositoryProvider) TransactionRepository(ctx context.Context) internal.ITransactionRepository {
	return NewTransactionRepository(MustGetTx(ctx))
}
