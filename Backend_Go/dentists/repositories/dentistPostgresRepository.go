package repositories

import (
	"segmentation/dentists/entities"
	dentistError "segmentation/dentists/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type dentistPosgresRepository struct {
	db *gorm.DB
}

func NewDentistPostgresRepository(db *gorm.DB) DentistRepository {
	return &dentistPosgresRepository{db: db}
}

func (r *dentistPosgresRepository) Search(key string, value *string) (bool, error) {
	dentist := new(entities.Dentist)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(dentist)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &dentistError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *dentistPosgresRepository) GetDentistDataByKey(key string, value *string) (*entities.Dentist, error) {
	dentist := new(entities.Dentist)
	dentist_data := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(dentist)
	if dentist_data.Error != nil {
		return nil, &dentistError.ServerInternalError{Err: dentist_data.Error}
	}
	return dentist, nil
}

func (r *dentistPosgresRepository) InsertDentistData(in *entities.InsertDentist) error {
	data := &entities.Dentist{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Password:  in.Password,
		Status:    in.Status,
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertDentistData:%v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *dentistPosgresRepository) UpdateDentistData(in *entities.UpdateDentist, id *uint32) error {
	data := &entities.UpdateDentist{
		FirstName:    in.FirstName,
		LastName:     in.LastName,
		Gender:       in.Gender,
		StartDate:    in.StartDate,
		YearExp:      in.YearExp,
		HospitalName: in.HospitalName,
		Phone:        in.Phone,
		Status:       in.Status,
	}
	result := r.db.Model(&entities.Dentist{}).Where("id = ?", *id).Updates(data)
	if result.Error != nil {
		log.Errorf("UpdateDentistData:%v", result.Error)
		return result.Error
	}
	log.Debugf("UpdateDentistData: %v", result.RowsAffected)
	return nil
}
