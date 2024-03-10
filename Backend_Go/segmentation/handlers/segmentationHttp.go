package handlers

import (
	"segmentation/segmentation/usecases"

	"github.com/labstack/echo/v4"
)

type segmentationHttpHandler struct {
	resultUsecase usecases.ResultUsecase
}

func NewSegmentationHttpHandler(resultUsecase usecases.ResultUsecase) SegmentationHandler {
	return &segmentationHttpHandler{
		resultUsecase: resultUsecase,
	}
}

func (s *segmentationHttpHandler) SaveSegmentation(c echo.Context) error {
	return nil
}
