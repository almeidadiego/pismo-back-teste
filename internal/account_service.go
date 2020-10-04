package internal

import (
	"context"
	"pismo-back-teste/internal/api/dto"
)

type AccountService struct {
	repo IAccountRepository
}

func NewAccountService(r IAccountRepository) *AccountService {
	return &AccountService{repo: r}
}

func (s *AccountService) GetAccount(ctx context.Context, id int) (dto.Account, error) {
	return s.repo.GetAccount(ctx, id)
}

func (s *AccountService) CreateAccount(ctx context.Context, account dto.Account) error {
	return s.repo.CreateAccount(ctx, account)
}
