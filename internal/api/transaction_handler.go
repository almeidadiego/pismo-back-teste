package api

import (
	"pismo-back-teste/internal"
	"pismo-back-teste/internal/api/dto"

	"github.com/labstack/echo"
)

type TransactionHandler struct {
	serviceProvider *internal.ServiceProvider
}

func NewTransactionHandler(services *internal.ServiceProvider) *TransactionHandler {
	return &TransactionHandler{serviceProvider: services}
}

func (h *TransactionHandler) Mount(v1 *echo.Group) {
	v1.POST("/transactions", h.post)
}

func (h *TransactionHandler) post(c echo.Context) error {
	ctx := c.Request().Context()

	transaction := dto.Transaction{}
	if err := c.Bind(&transaction); err != nil {
		return err
	}

	service := h.serviceProvider.TransactionService(ctx)
	err := service.CreateTransaction(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}
