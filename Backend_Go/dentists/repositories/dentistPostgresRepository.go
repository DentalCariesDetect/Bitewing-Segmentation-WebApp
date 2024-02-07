package repositories

import (
	"segmentation/dentists/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type dentistPosgresRepository struct {
	db *gorm.DB
}

func NewDentistPostgresRepository(db *gorm.DB) DentistRepository {
	return &dentistPosgresRepository{db: db}
}

func (r *dentistPosgresRepository) InsertDentistData(in *entities.InsertDentist) error {
	data := &entities.Dentist{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Password:  in.Password,
		CreateOn:  in.CreateOn,
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
