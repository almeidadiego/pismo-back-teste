package internal

import (
	"context"
	"pismo-back-teste/internal/api/dto"
)

type IRepositoryProvider interface {
	AccountRepository(ctx context.Context) IAccountRepository
}

type IAccountRepository interface {
	GetAccount(ctx context.Context, id int) (dto.Account, error)
}
