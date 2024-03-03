package handlers

import "github.com/labstack/echo/v4"

type DentistHandler interface {
	UpdateDentist(c echo.Context) error
}
