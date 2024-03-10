package server

import (
	"net/http"
	tokenUsecase "segmentation/auth/usecases"
	"segmentation/configs"
	dentistRepository "segmentation/dentists/repositories"
	"strconv"

	"github.com/labstack/echo/v4"
)

func TokenAuthentication(repo dentistRepository.DentistRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, "missing authorization header")
			}
			// check if token not start with Bearer
			if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
				return c.JSON(http.StatusUnauthorized, "invalid or expired token")
			} else {
				authHeader = authHeader[7:]
			}

			tokenUsecase := tokenUsecase.NewTokenUsecaseImpl(configs.GetJwtConfig().SecretKey)
			userID, err := tokenUsecase.ParseToken(&authHeader)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "invalid or expired token")
			}
			dentistId := strconv.FormatUint(uint64(*userID), 10)
			if result, err := repo.Search("id", &dentistId); !result || err != nil {
				return c.JSON(http.StatusUnauthorized, "invalid or expired token")
			}
			return next(c)
		}
	}
}
