package repositories

import "segmentation/patient/entities"

type PatientRepository interface {
	Search(key string, value *string) (bool, error)
	GetPatientDataByKey(key string, value *string) (*entities.Patient, error)
	GetPatientsDataByKey(key string, value *string) ([]*entities.Patient, error)
	InsertPatientData(in *entities.Patient) error
	UpdatePatientData(in *entities.Patient) error
	DeletePatientData(id *uint64) error
}
