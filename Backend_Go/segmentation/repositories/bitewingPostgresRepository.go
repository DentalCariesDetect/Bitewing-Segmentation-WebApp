package repositories

import (
	"segmentation/segmentation/entities"
	segmentationError "segmentation/segmentation/errors"

	"time"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type bitewingPosgresRepository struct {
	db *gorm.DB
}

func NewBitewingPostgresRepository(db *gorm.DB) BitewingRepository {
	return &bitewingPosgresRepository{db: db}
}

func (r *bitewingPosgresRepository) Search(key string, value *string) (bool, error) {
	bitewing := new(entities.Bitewing)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(bitewing)
	if result.RowsAffected > 0 {
		return true, nil
	} else {
		if result.Error != nil {
			return false, &segmentationError.ServerInternalError{Err: result.Error}
		} else {
			return false, nil
		}
	}
}

func (r *bitewingPosgresRepository) GetBitewingDataByKey(key string, value *string) (*entities.Bitewing, error) {
	bitewing := new(entities.Bitewing)
	dentist_data := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(bitewing)
	if dentist_data.Error != nil {
		return nil, &segmentationError.ServerInternalError{Err: dentist_data.Error}
	}
	return bitewing, nil
}

func (r *bitewingPosgresRepository) InsertBitewingData(in *entities.Bitewing) error {
	result := r.db.Create(in)

	if result.Error != nil {
		log.Errorf("InsertDentistData:%v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *bitewingPosgresRepository) UpdateBitewingData(in *entities.Bitewing) error {
	result := r.db.Save(in)

	if result.Error != nil {
		log.Errorf("UpdateDentistData:%v", result.Error)
		return result.Error
	}

	log.Debugf("UpdateCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *bitewingPosgresRepository) DeleteBitewingData(id *uint64) error {
	result := r.db.Model(&entities.Bitewing{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
		"status":     "Removed",
		"deleted_at": time.Now(),
	})
	if result.Error != nil {
		log.Errorf("DeleteDentistData:%v", result.Error)
		return result.Error
	}

	log.Debugf("DeleteCockroachData: %v", result.RowsAffected)
	return nil
}
