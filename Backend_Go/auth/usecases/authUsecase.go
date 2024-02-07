package usecases

import (
	"segmentation/auth/models"
	"segmentation/dentists/entities"
	"segmentation/dentists/repositories"
	"time"
)

type AuthUsecase interface {
	RegisterDataProcessing(in *models.RegisterData) error
}

type authUsecaseImpl struct {
	dentistRepository repositories.DentistRepository
}

func NewAuthUsecaseImpl(dentistRepository repositories.DentistRepository) AuthUsecase {
	return &authUsecaseImpl{
		dentistRepository: dentistRepository,
	}
}

func (u *authUsecaseImpl) RegisterDataProcessing(in *models.RegisterData) error {
	insertDentistData := &entities.InsertDentist{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Password:  in.Password,
		CreateOn:  time.Now(),
		Status:    "Active",
	}

	if err := u.dentistRepository.InsertDentistData(insertDentistData); err != nil {
		return err
	}

	return nil
}
