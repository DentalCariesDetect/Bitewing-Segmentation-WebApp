package usecases

import (
	"segmentation/dentists/entities"
	dentistError "segmentation/dentists/errors"
	"segmentation/dentists/models"
	"segmentation/dentists/repositories"
)

type DentistUsecase interface {
	UpdateDentist(in *models.UpdateModel, id uint32) error
	CheckDentistID(id *string) error
}

type dentistUsecaseImpl struct {
	dentistRepository repositories.DentistRepository
}

func NewDentistUsecaseImpl(dentistRepository repositories.DentistRepository) DentistUsecase {
	return &dentistUsecaseImpl{
		dentistRepository: dentistRepository,
	}
}

func (u *dentistUsecaseImpl) CheckDentistID(id *string) error {
	if result, err := u.dentistRepository.Search("id", id); !result || err != nil {
		if err != nil {
			return &dentistError.ServerInternalError{Err: err}
		}
		return &dentistError.DentistNotFoundError{}
	}
	return nil
}

func (u *dentistUsecaseImpl) UpdateDentist(in *models.UpdateModel, id uint32) error {
	updateDentistData := &entities.UpdateDentist{
		FirstName:    in.FirstName,
		LastName:     in.LastName,
		Gender:       &in.Gender,
		HospitalName: &in.HospitalName,
		Phone:        &in.Phone,
		StartDate:    &in.StartDate,
		YearExp:      &in.YearExp,
		Status:       "Active",
	}
	if err := u.dentistRepository.UpdateDentistData(updateDentistData, &id); err != nil {
		return err
	}
	return nil
}
