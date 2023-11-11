package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pallavi23s/Go-Echo-Api/cmd/api/handelers"
)

func main() {
	e := echo.New()
	e.GET("/health-check", handelers.HealthCheckHandeler)
	e.GET("/posts", handelers.PostIndexHandeler)
	e.GET("/post/:id", handelers.PostSingleHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
