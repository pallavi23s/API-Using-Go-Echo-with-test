package handelers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pallavi23s/Go-Echo-Api/cmd/api/service"
)

func PostIndexHandeler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)

	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}

	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)

}
