package v1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sula7/network-info/models"
)

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func addresses(c echo.Context) error {
	ipAddresses := models.IPAddresses{}
	err := c.Bind(&ipAddresses.List)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	for _, ip := range ipAddresses.List {
		fmt.Println(ip.Address)
	}
	return c.String(http.StatusOK, "OK")
}
