package api

import (
	"database/sql"
	"pismo-back-teste/internal"
	"pismo-back-teste/internal/api/middleware"

	"github.com/labstack/echo"
)

type Options struct {
	ServiceProvider  *internal.ServiceProvider
	DatabaseProvider *sql.DB
}

// InitRoutes sets up routing.
func InitRoutes(e *echo.Echo, opts Options) {
	// Setup sql.Tx on context.
	e.Use(middleware.RequestTransaction(opts.DatabaseProvider))

	api := e.Group("/api/v1")
	NewAccountHandler(opts.ServiceProvider).Mount(api)
	NewTransactionHandler(opts.ServiceProvider).Mount(api)
}
