package api

import (
	"net/http"
	"pismo-back-teste/internal"
	"strconv"

	"github.com/labstack/echo"
)

type AccountHandler struct {
	serviceProvider *internal.ServiceProvider
}

func NewAccountHandler(services *internal.ServiceProvider) *AccountHandler {
	return &AccountHandler{serviceProvider: services}
}

func (h *AccountHandler) Mount(v1 *echo.Group) {
	v1.GET("/accounts/:accountId", h.get)
}

func (h *AccountHandler) get(c echo.Context) error {
	ctx := c.Request().Context()

	accountID, err := strconv.Atoi(c.Param("accountId"))
	if err != nil {
		return err
	}

	service := h.serviceProvider.AccountService(ctx)
	account, err := service.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}

	httpStatus := http.StatusOK
	if account.ID < 1 {
		httpStatus = http.StatusNotFound
	}

	return c.JSON(httpStatus, account)
}
