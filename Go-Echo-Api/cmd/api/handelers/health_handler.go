package handelers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheckHandeler(c echo.Context) error {
	return c.String(http.StatusOK, "OK!")
}
