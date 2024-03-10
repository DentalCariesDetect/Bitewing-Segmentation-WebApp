package handlers

import "github.com/labstack/echo/v4"

type SegmentationHandler interface {
	SaveSegmentation(c echo.Context) error
}
