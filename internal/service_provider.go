package internal

import "context"

type ServiceProvider struct {
	repoProvider IRepositoryProvider
}

func NewServiceProvider(r IRepositoryProvider) *ServiceProvider {
	return &ServiceProvider{repoProvider: r}
}

func (s *ServiceProvider) AccountService(ctx context.Context) *AccountService {
	return NewAccountService(s.repoProvider.AccountRepository(ctx))
}

func (s *ServiceProvider) TransactionService(ctx context.Context) *TransactionService {
	return NewTransactionService(s.repoProvider.TransactionRepository(ctx))
}
