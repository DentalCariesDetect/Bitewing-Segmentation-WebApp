package handlers

import "github.com/labstack/echo/v4"

type PatientHandler interface {
	CreatePatient(c echo.Context) error
	GetPatientById(c echo.Context) error
	GetPatientsByDentist(c echo.Context) error
	UpdatePatient(c echo.Context) error
	DeletePatient(c echo.Context) error
}
