package handlers

import "github.com/labstack/echo/v4"

type SegmentationHandler interface {
	SaveResult(c echo.Context) error
	GetDetailResult(c echo.Context) error
	GetResults(c echo.Context) error
	DeleteResult(c echo.Context) error
}
