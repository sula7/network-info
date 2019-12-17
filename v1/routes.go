package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewApi(bindPort string) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/ping", ping)

	g := e.Group("/api/v1")
	g.POST("/addresses", addresses)

	e.Logger.Fatal(e.Start(bindPort))
}
