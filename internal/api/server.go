package api

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func NewServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		log.Error(err)
	}
	return e
}
