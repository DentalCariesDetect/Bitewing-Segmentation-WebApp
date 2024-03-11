package repositories

import (
	"segmentation/segmentation/entities"
	segmentationError "segmentation/segmentation/errors"
	"time"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type toothPosgresRepository struct {
	db *gorm.DB
}

func NewToothPostgresRepository(db *gorm.DB) ToothRepository {
	return &toothPosgresRepository{db: db}
}

func (r *toothPosgresRepository) Search(key string, value *string) (bool, error) {
	tooth := new(entities.Tooth)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(tooth)
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
func (r *toothPosgresRepository) SearchWithResultId(key string, value *string, resultId *string) (bool, error) {
	tooth := new(entities.Tooth)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Where("result_id = ?", *resultId).Limit(1).Find(tooth)
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

func (r *toothPosgresRepository) GetToothDataByKey(key string, value *string) (*entities.Tooth, error) {
	tooth := new(entities.Tooth)
	dentist_data := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Preload("Image").First(tooth)
	if dentist_data.Error != nil {
		return nil, &segmentationError.ServerInternalError{Err: dentist_data.Error}
	}
	return tooth, nil
}
func (r *toothPosgresRepository) GetTeethDataByKey(key string, value *string) ([]*entities.Tooth, error) {
	tooth := []*entities.Tooth{}
	dentist_data := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Preload("Image").Find(&tooth)
	if dentist_data.Error != nil {
		return nil, &segmentationError.ServerInternalError{Err: dentist_data.Error}
	}
	return tooth, nil
}

func (r *toothPosgresRepository) InsertToothData(in *entities.Tooth) error {
	result := r.db.Create(in)

	if result.Error != nil {
		log.Errorf("InsertDentistData:%v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *toothPosgresRepository) UpdateToothData(in *entities.Tooth) error {
	result := r.db.Save(in)

	if result.Error != nil {
		log.Errorf("UpdateDentistData:%v", result.Error)
		return result.Error
	}

	log.Debugf("UpdateCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *toothPosgresRepository) DeleteToothData(id *uint64) error {
	result := r.db.Model(&entities.Tooth{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
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
