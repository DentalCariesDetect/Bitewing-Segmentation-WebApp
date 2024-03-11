package usecases

import (
	"fmt"
	"segmentation/patient/entities"
	patientError "segmentation/patient/errors"
	"segmentation/patient/models"
	"segmentation/patient/repositories"
	"strconv"
)

type PatientUsecase interface {
	CheckPatientId(id *string) error
	CreatePatient(in *models.PatientInsertModel, id *string) error
	GetPatientById(id *string) (*models.PatientModel, error)
	GetPatientsByDentist(dentistId *string) ([]*models.PatientModel, error)
	UpdatePatient(in *models.PatientUpdateModel, id *string) error
	DeletePatient(id *string) error
}

type patientUsecaseImpl struct {
	patientRepository repositories.PatientRepository
}

func NewPatientUsecaseImpl(patientRepository repositories.PatientRepository) PatientUsecase {
	return &patientUsecaseImpl{
		patientRepository: patientRepository,
	}
}

func (u *patientUsecaseImpl) CheckPatientId(id *string) error {
	if result, err := u.patientRepository.Search("id", id); !result || err != nil {
		if err != nil {
			return err
		} else {
			return &patientError.PatientNotFoundError{}
		}
	}
	return nil
}

func (u *patientUsecaseImpl) GetPatientById(id *string) (*models.PatientModel, error) {
	patientData, err := u.patientRepository.GetPatientDataByKey("id", id)
	if err != nil {
		return nil, err
	}
	patientModel := &models.PatientModel{
		ID:        uint64(patientData.ID),
		BirthDate: patientData.BirthDate,
		Gender:    patientData.Gender,
		Age:       patientData.Age,
		Phone:     patientData.Phone,
	}
	return patientModel, nil
}

func (u *patientUsecaseImpl) GetPatientsByDentist(dentistId *string) ([]*models.PatientModel, error) {
	patients, err := u.patientRepository.GetPatientsDataByKey("dentist_id", dentistId)
	if err != nil {
		return nil, err
	}
	patientModels := make([]*models.PatientModel, 0)
	for _, patient := range patients {
		patientModel := &models.PatientModel{
			ID:        uint64(patient.ID),
			BirthDate: patient.BirthDate,
			Gender:    patient.Gender,
			Age:       patient.Age,
			Phone:     patient.Phone,
		}
		patientModels = append(patientModels, patientModel)
	}
	return patientModels, nil
}

func (u *patientUsecaseImpl) CreatePatient(in *models.PatientInsertModel, id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &patientError.ServerInternalError{}
	}
	patient := &entities.Patient{
		DentistID: uint(idUint64),
		BirthDate: in.BirthDate,
		Gender:    in.Gender,
		Age:       in.Age,
		Phone:     stringIsNil(in.Phone),
		Status:    "Active",
	}
	if err := u.patientRepository.InsertPatientData(patient); err != nil {
		return err
	}
	return nil
}

func (u *patientUsecaseImpl) UpdatePatient(in *models.PatientUpdateModel, id *string) error {
	patient, err := u.patientRepository.GetPatientDataByKey("id", id)
	fmt.Println(patient)
	if err != nil {
		return err
	}
	if patient == nil {
		return &patientError.PatientNotFoundError{}
	}
	patient.BirthDate = in.BirthDate
	patient.Gender = in.Gender
	patient.Age = in.Age
	patient.Phone = stringIsNil(in.Phone)

	if err := u.patientRepository.UpdatePatientData(patient); err != nil {
		return err
	}
	return nil
}

func (u *patientUsecaseImpl) DeletePatient(id *string) error {
	idUint64, err := strconv.ParseUint(*id, 10, 64)
	if err != nil {
		return &patientError.ServerInternalError{}
	}
	if err := u.patientRepository.DeletePatientData(&idUint64); err != nil {
		return err
	}
	return nil
}
