package usecases

import (
	"segmentation/dentists/entities"
	dentistError "segmentation/dentists/errors"
	"segmentation/dentists/models"
	"segmentation/dentists/repositories"
	"strconv"
)

type DentistUsecase interface {
	GetDentistById(id *string) (*models.DentistModel, error)
	UpdateDentist(in *models.UpdateModel, id *string) error
	CheckDentistId(id *string) error
	GetDentistAll() ([]*models.DentistModel, error)
	DeleteDentist(id *string) error
}

type dentistUsecaseImpl struct {
	dentistRepository repositories.DentistRepository
}

func NewDentistUsecaseImpl(dentistRepository repositories.DentistRepository) DentistUsecase {
	return &dentistUsecaseImpl{
		dentistRepository: dentistRepository,
	}
}

func (u *dentistUsecaseImpl) CheckDentistId(id *string) error {
	if result, err := u.dentistRepository.Search("id", id); !result || err != nil {
		if err != nil {
			return &dentistError.ServerInternalError{Err: err}
		}
		return &dentistError.DentistNotFoundError{}
	}
	return nil
}

func (u *dentistUsecaseImpl) GetDentistById(id *string) (*models.DentistModel, error) {
	dentistData, err := u.dentistRepository.GetDentistDataByKey("id", id)
	if err != nil {
		return nil, err
	}
	dentistModel := &models.DentistModel{
		ID:           uint64(dentistData.ID),
		FirstName:    dentistData.FirstName,
		LastName:     dentistData.LastName,
		Username:     dentistData.Username,
		HospitalName: dentistData.HospitalName,
		Phone:        dentistData.Phone,
		Gender:       dentistData.Gender,
		StartDate:    dentistData.StartDate,
		YearExp:      dentistData.YearExp,
	}
	return dentistModel, nil
}

func (u *dentistUsecaseImpl) GetDentistAll() ([]*models.DentistModel, error) {
	dentists, err := u.dentistRepository.GetDentistDataAll()
	if err != nil {
		return nil, err
	}
	dentistModels := []*models.DentistModel{}
	for _, dentist := range dentists {
		dentistModel := &models.DentistModel{
			ID:           uint64(dentist.ID),
			FirstName:    dentist.FirstName,
			LastName:     dentist.LastName,
			Username:     dentist.Username,
			HospitalName: dentist.HospitalName,
			Phone:        dentist.Phone,
			Gender:       dentist.Gender,
			StartDate:    dentist.StartDate,
			YearExp:      dentist.YearExp,
		}
		dentistModels = append(dentistModels, dentistModel)
	}
	return dentistModels, nil
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

func (u *dentistUsecaseImpl) DeleteDentist(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &dentistError.ServerInternalError{Err: err}
	}
	if err := u.dentistRepository.DeleteDentistData(&idUint64); err != nil {
		return err
	}
	return nil
}
