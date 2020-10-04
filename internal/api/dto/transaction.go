package dto

import "time"

type OperationType int

const (
	SpotPurchase OperationType = 1 + iota
	HirePurchase
	Withdraw
	Payment
)

var operationTypes = []string{
	"SPOT_PURCHASE",
	"HIRE_PURCHASE",
	"WITHDRAW",
	"PAYMENT",
}

func (t OperationType) String() string {
	return operationTypes[t-1]
}

type Transaction struct {
	ID              int           `json:"id"`
	AccountID       int           `json:"account_id"`
	OperationTypeID OperationType `json:"operation_type_id"`
	Amount          float64       `json:"amount"`
	CreatedAt       time.Time     `json:"created_at"`
}
