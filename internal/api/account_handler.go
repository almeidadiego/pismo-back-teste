package api

import (
	"net/http"
	"pismo-back-teste/internal"
	"pismo-back-teste/internal/api/dto"
	"strconv"

	"github.com/labstack/echo"
)

type AccountHandler struct {
	serviceProvider *internal.ServiceProvider
}

func NewAccountHandler(serviceProvider *internal.ServiceProvider) *AccountHandler {
	return &AccountHandler{serviceProvider: serviceProvider}
}

func (h *AccountHandler) Mount(v1 *echo.Group) {
	v1.GET("/accounts/:accountId", h.get)
	v1.POST("/accounts", h.post)
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

func (h *AccountHandler) post(c echo.Context) error {
	ctx := c.Request().Context()

	account := dto.Account{}
	if err := c.Bind(&account); err != nil {
		return err
	}

	service := h.serviceProvider.AccountService(ctx)
	if err := service.CreateAccount(ctx, account); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, account)
}
