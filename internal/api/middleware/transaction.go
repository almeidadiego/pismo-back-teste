package middleware

import (
	"database/sql"
	"fmt"
	"pismo-back-teste/internal/database"

	"github.com/labstack/echo"
)

// RequestTransaction is a Middleware function chained in the HTTP request-response
// cycle with access to Echo#Context which it uses to perform a specific action. In
// this case, this function creates a db transaction and adds it to Echo#Context
func RequestTransaction(db *sql.DB) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			ctx := r.Context()

			tx, err := db.BeginTx(ctx, nil)
			if err != nil {
				fmt.Println(err)
				return err
			}

			// ensure we finish tx in case of panic
			defer tx.Rollback() // nolint: errcheck

			ctx = database.WithTx(ctx, tx)
			c.SetRequest(r.WithContext(ctx))

			nextErr := next(c)
			if nextErr != nil {
				if txErr := tx.Rollback(); txErr != nil {
					// don't want to overwrite err here, just log it
					fmt.Println(err)
				}
				return nextErr
			}

			if err := tx.Commit(); err != nil {
				return err
			}

			return nil
		}
	}
}
