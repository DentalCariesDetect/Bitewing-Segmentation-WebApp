package handlers

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	Register(c echo.Context) error
	//Login(c echo.Context) error
}
