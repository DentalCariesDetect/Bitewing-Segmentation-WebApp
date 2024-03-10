package usecases

import (
	"segmentation/segmentation/repositories"

	"github.com/labstack/echo/v4"
)

type ResultUsecase interface {
	SaveResult(c echo.Context) error
}

type resultUsecaseImpl struct {
	resultRepository   repositories.ResultRepository
	bitewingRepository repositories.BitewingRepository
	toothRepository    repositories.ToothRepository
}

func NewResultUsecaseImpl(resultRepository repositories.ResultRepository, bitewingRepository repositories.BitewingRepository, toothRepository repositories.ToothRepository) ResultUsecase {
	return &resultUsecaseImpl{
		resultRepository:   resultRepository,
		bitewingRepository: bitewingRepository,
		toothRepository:    toothRepository,
	}
}

func (r *resultUsecaseImpl) SaveResult(c echo.Context) error {
	return nil
}
