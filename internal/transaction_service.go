package internal

import (
	"context"
	"math"
	"pismo-back-teste/internal/api/dto"
)

type TransactionService struct {
	repo ITransactionRepository
}

func NewTransactionService(r ITransactionRepository) *TransactionService {
	return &TransactionService{repo: r}
}

// CreateTransaction saves a new transaction
func (s *TransactionService) CreateTransaction(
	ctx context.Context, transaction dto.Transaction) error {
	checkOperationType(&transaction)
	return s.repo.CreateTransaction(ctx, transaction)
}

func checkOperationType(t *dto.Transaction) {
	t.Amount = math.Abs(t.Amount)
	if t.OperationTypeID != dto.Payment {
		t.Amount = -t.Amount
	}
}
