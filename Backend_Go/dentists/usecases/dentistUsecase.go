package usecases

import (
	"segmentation/dentists/entities"
	dentistError "segmentation/dentists/errors"
	"segmentation/dentists/models"
	"segmentation/dentists/repositories"
	"strconv"
)

type DentistUsecase interface {
	UpdateDentist(in *models.UpdateModel, id *string) error
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

func (u *dentistUsecaseImpl) UpdateDentist(in *models.UpdateModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &dentistError.ServerInternalError{Err: err}
	}
	updateDentistData := &entities.UpdateDentist{
		FirstName:    in.FirstName,
		LastName:     in.LastName,
		Gender:       stringIsNil(in.Gender),
		HospitalName: stringIsNil(in.HospitalName),
		Phone:        stringIsNil(in.Phone),
		StartDate:    stringIsNil(in.StartDate),
		YearExp:      uintIsNil(in.YearExp),
		Status:       "Active",
	}
	if err := u.dentistRepository.UpdateDentistData(updateDentistData, &idUint64); err != nil {
		return err
	}
	return nil
}
