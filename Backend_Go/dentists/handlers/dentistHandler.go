package handlers

import "github.com/labstack/echo/v4"

type DentistHandler interface {
	UpdateDentist(c echo.Context) error
	GetDentistById(c echo.Context) error
	GetDentistAll(c echo.Context) error
	DeleteDentist(c echo.Context) error
}
