package internal

import "context"

// ServiceProvider creates services scoped to the given context
type ServiceProvider struct {
	repoProvider IRepositoryProvider
}

func NewServiceProvider(r IRepositoryProvider) *ServiceProvider {
	return &ServiceProvider{repoProvider: r}
}

// AccountService returns an account service
func (s *ServiceProvider) AccountService(ctx context.Context) *AccountService {
	return NewAccountService(s.repoProvider.AccountRepository(ctx))
}

// TransactionService returns an transaction service
func (s *ServiceProvider) TransactionService(ctx context.Context) *TransactionService {
	return NewTransactionService(s.repoProvider.TransactionRepository(ctx))
}
