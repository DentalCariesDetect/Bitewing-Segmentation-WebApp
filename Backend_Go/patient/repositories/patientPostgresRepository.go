package repositories

import (
	"segmentation/patient/entities"
	patientError "segmentation/patient/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type patientPosgresRepository struct {
	db *gorm.DB
}

func NewPatientPostgresRepository(db *gorm.DB) PatientRepository {
	return &patientPosgresRepository{db: db}
}

func (r *patientPosgresRepository) Search(key string, value *string) (bool, error) {
	patient := new(entities.Patient)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(patient)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &patientError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *patientPosgresRepository) GetPatientDataByKey(key string, value *string) (*entities.Patient, error) {
	patient := new(entities.Patient)
	patient_data := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(patient)
	if patient_data.Error != nil {
		return nil, &patientError.ServerInternalError{Err: patient_data.Error}
	}
	return patient, nil
}
func (r *patientPosgresRepository) GetPatientsDataByKey(key string, value *string) ([]*entities.Patient, error) {
	patients := []*entities.Patient{}
	patient_datas := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Find(&patients)
	if patient_datas.Error != nil {
		return nil, &patientError.ServerInternalError{Err: patient_datas.Error}
	}
	return patients, nil
}

func (r *patientPosgresRepository) InsertPatientData(in *entities.Patient) error {
	result := r.db.Create(in)

	if result.Error != nil {
		log.Errorf("InsertpatientData:%v", result.Error)
		return &patientError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *patientPosgresRepository) UpdatePatientData(in *entities.Patient) error {
	result := r.db.Save(in)
	if result.Error != nil {
		log.Errorf("UpdatePatientData:%v", result.Error)
		return &patientError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("UpdatePatientData: %v", result.RowsAffected)
	return nil
}

func (r *patientPosgresRepository) DeletePatientData(id *uint64) error {
	result := r.db.Model(&entities.Patient{}).Where("id = ?", *id).Where("status <> ?", "Removed").Update("status", "Removed")
	if result.Error != nil {
		log.Errorf("DeletePatientData:%v", result.Error)
		return &patientError.ServerInternalError{Err: result.Error}
	}
	log.Debugf("DeletePatientData: %v", result.RowsAffected)
	return nil
}
