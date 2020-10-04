package internal

import (
	"context"
	"errors"
	"pismo-back-teste/internal/api/dto"
)

type TransactionService struct {
	repo ITransactionRepository
}

func NewTransactionService(r ITransactionRepository) *TransactionService {
	return &TransactionService{repo: r}
}

func (s *TransactionService) CreateTransaction(
	ctx context.Context, transaction dto.Transaction) error {
	if !correctOperationSign(transaction) {
		return errors.New(
			"Wrong operation sign for operation type " + transaction.OperationTypeID.String(),
		)
	}

	return s.repo.CreateTransaction(ctx, transaction)
}

func correctOperationSign(t dto.Transaction) bool {
	if t.OperationTypeID == dto.Payment {
		return t.Amount > 0
	}

	return t.Amount < 0
}
