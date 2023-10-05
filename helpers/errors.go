package helpers

import (
	"github.com/labstack/echo/v4"
)

func ErrorWrap(err error, code int) error {
	return echo.NewHTTPError(code, err)
}
