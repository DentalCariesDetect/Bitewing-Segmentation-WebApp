package handlers

import (
	"fmt"
	"net/http"
	patientError "segmentation/patient/errors"
	"segmentation/patient/models"
	"segmentation/patient/usecases"

	"github.com/gofiber/fiber/v2/log"
	"github.com/labstack/echo/v4"
)

type patientHttpHandler struct {
	patientUsecase usecases.PatientUsecase
}

func NewPatientHttpHandler(patientUsecase usecases.PatientUsecase) PatientHandler {
	return &patientHttpHandler{
		patientUsecase: patientUsecase,
	}
}

func (h *patientHttpHandler) GetPatientById(c echo.Context) error {
	id := c.Param("id")
	if err := h.patientUsecase.CheckPatientId(&id); err != nil {
		if _, ok := err.(*patientError.PatientNotFoundError); ok {
			return response(c, 404, "Patient not found", nil)
		} else {
			return response(c, 500, "Server Internal Error", nil)
		}
	}
	patient, err := h.patientUsecase.GetPatientById(&id)
	if err != nil {
		return response(c, 500, "Server Internal Error", nil)
	}
	return response(c, 200, "Patient found", patient)
}

func (h *patientHttpHandler) GetPatientsByDentist(c echo.Context) error {
	dentistId := c.Get("dentistId").(string)
	patients, err := h.patientUsecase.GetPatientsByDentist(&dentistId)
	if err != nil {
		return response(c, 500, "Server Internal Error", nil)
	}
	return response(c, 200, "Patients found", patients)
}

func (h *patientHttpHandler) CreatePatient(c echo.Context) error {
	reqBody := new(models.PatientInsertModel)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Invalid request body", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	dentistId := c.Get("dentistId").(string)
	if err := h.patientUsecase.CreatePatient(reqBody, &dentistId); err != nil {
		if _, ok := err.(*patientError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 201, "Patient created", nil)
}

func (h *patientHttpHandler) UpdatePatient(c echo.Context) error {
	patientId := c.Param("id")
	reqBody := new(models.PatientUpdateModel)

	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Invalid request body", nil)
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return validationErrorResponse(c, err)
	}
	if err := h.patientUsecase.CheckPatientId(&patientId); err != nil {
		if _, ok := err.(*patientError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	fmt.Println("patientId: ", patientId)
	if err := h.patientUsecase.UpdatePatient(reqBody, &patientId); err != nil {
		if _, ok := err.(*patientError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Patient updated", nil)
}

func (h *patientHttpHandler) DeletePatient(c echo.Context) error {
	patientId := c.Param("id")
	if err := h.patientUsecase.CheckPatientId(&patientId); err != nil {
		if _, ok := err.(*patientError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	if err := h.patientUsecase.DeletePatient(&patientId); err != nil {
		if _, ok := err.(*patientError.ServerInternalError); ok {
			return response(c, 500, "Server Internal Error", nil)
		} else {
			return response(c, 400, err.Error(), nil)
		}
	}
	return response(c, 200, "Patient deleted", nil)
}
