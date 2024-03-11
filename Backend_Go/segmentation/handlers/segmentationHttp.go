package handlers

import (
	segmentationError "segmentation/segmentation/errors"
	"segmentation/segmentation/models"
	"segmentation/segmentation/usecases"

	"github.com/gofiber/fiber/v2/log"
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

func (h *segmentationHttpHandler) GetDetailResult(c echo.Context) error {
	resultId := c.Param("id")
	if err := h.resultUsecase.CheckResultId(&resultId); err != nil {
		if _, ok := err.(*segmentationError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	result, err := h.resultUsecase.GetDetailResult(&resultId)
	if err != nil {
		return response(c, 500, "Server Internal Error", nil)
	}
	return response(c, 200, "Success", result)

}

func (h *segmentationHttpHandler) SaveResult(c echo.Context) error {
	resultId := c.Param("id")
	if err := h.resultUsecase.CheckResultId(&resultId); err != nil {
		if _, ok := err.(*segmentationError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	reqBody := new(models.SegmentationUpdateModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return response(c, 400, "Bad Request", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.resultUsecase.SaveResult(reqBody, &resultId); err != nil {
		if _, ok := err.(*segmentationError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}

func (h *segmentationHttpHandler) GetResults(c echo.Context) error {
	dentistId := c.Get("dentistId").(string)
	results, err := h.resultUsecase.GetResultsByDentistId(&dentistId)
	if err != nil {
		if _, ok := err.(*segmentationError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", results)
}

func (h *segmentationHttpHandler) DeleteResult(c echo.Context) error {
	resultId := c.Param("id")
	if err := h.resultUsecase.CheckResultId(&resultId); err != nil {
		if _, ok := err.(*segmentationError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	if err := h.resultUsecase.DeleteResult(&resultId); err != nil {
		if _, ok := err.(*segmentationError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Success", nil)
}
