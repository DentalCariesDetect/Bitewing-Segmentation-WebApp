package usecases

import (
	"errors"
	auth_error "segmentation/auth/entities"
	"segmentation/auth/models"
	"segmentation/dentists/entities"
	"segmentation/dentists/repositories"
	"time"
)

type AuthUsecase interface {
	RegisterDataProcessing(in *models.RegisterData) error
	CheckData(in *models.RegisterData) error
}

type authUsecaseImpl struct {
	dentistRepository repositories.DentistRepository
}

func NewAuthUsecaseImpl(dentistRepository repositories.DentistRepository) AuthUsecase {
	return &authUsecaseImpl{
		dentistRepository: dentistRepository,
	}
}
func (u *authUsecaseImpl) CheckData(in *models.RegisterData) error {
	username := &in.Username
	//result = true -> found username
	//result = false -> Not found username
	//err != nill -> Found error
	if result, err := u.dentistRepository.SearchUsername(username); result || err != nil {
		return &auth_error.UsernameAlreadyExistError{}
	} else {
		if errors.Is(err, &auth_error.ServerInternalError{}) {
			return &auth_error.ServerInternalError{Err: err}
		}
	}
	return nil
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
		return &auth_error.ServerInternalError{Err: err}
	}

	return nil
}
