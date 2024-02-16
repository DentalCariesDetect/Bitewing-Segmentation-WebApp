package handlers

import (
	"net/http"
	auth_error "segmentation/auth/entities"
	"segmentation/auth/models"
	"segmentation/auth/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type authHttpHandler struct {
	authUsecase usecases.AuthUsecase
}

func NewAuthHttpHandler(authUsecase usecases.AuthUsecase) AuthHandler {
	return &authHttpHandler{
		authUsecase: authUsecase,
	}
}

func (h *authHttpHandler) Register(c echo.Context) error {
	reqBody := new(models.RegisterData)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}
	if err := h.authUsecase.CheckData(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*auth_error.ServerInternalError); ok {
			return response(c, http.StatusInternalServerError, "Server Internal Error")
		} else {
			return response(c, http.StatusBadRequest, err.Error())
		}
	}
	if err := h.authUsecase.RegisterDataProcessing(reqBody); err != nil {
		return response(c, http.StatusInternalServerError, "Processing Data failed")
	}

	return response(c, http.StatusOK, "Success")
}

func (h *authHttpHandler) Login(c echo.Context) error {
	reqBody := new(models.LoginData)
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}
	if err := c.Validate(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		return response(c, http.StatusBadRequest, "Bad request")
	}

	if token, err := h.authUsecase.LoginDataProcession(reqBody); err != nil {
		log.Errorf("Error validating request body: %v", err)
		if _, ok := err.(*auth_error.ServerInternalError); ok {
			return response(c, http.StatusInternalServerError, "Server Internal Error")
		} else {
			return response(c, http.StatusBadRequest, err.Error())
		}
	} else {
		return response(c, http.StatusOK, *token)
	}

}
