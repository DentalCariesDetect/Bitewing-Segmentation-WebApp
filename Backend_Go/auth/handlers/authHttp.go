package handlers

import (
	"fmt"
	"net/http"
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
	fmt.Println(reqBody)

	if err := h.authUsecase.RegisterDataProcessing(reqBody); err != nil {
		return response(c, http.StatusInternalServerError, "Processing Data failed")
	}

	return response(c, http.StatusOK, "Success")
}
