package handlers

import "github.com/labstack/echo"

type Employee interface {
	Number2(c echo.Context) error
}
