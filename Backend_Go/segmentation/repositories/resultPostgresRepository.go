package repositories

import (
	"segmentation/segmentation/entities"
	segmentationError "segmentation/segmentation/errors"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type resultPosgresRepository struct {
	db *gorm.DB
}

func NewResultPostgresRepository(db *gorm.DB) ResultRepository {
	return &resultPosgresRepository{db: db}
}

func (r *resultPosgresRepository) Search(key string, value *string) (bool, error) {
	result := new(entities.Result)
	result_t := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(result)
	if result_t.RowsAffected > 0 {
		return true, nil
	} else {
		if result_t.Error != nil {
			return false, &segmentationError.ServerInternalError{Err: result_t.Error}
		} else {
			return false, nil
		}
	}
}

func (r *resultPosgresRepository) GetResultDataByKey(key string, value *string) (*entities.Result, error) {
	result := new(entities.Result)
	result_t := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(result)
	if result_t.Error != nil {
		return nil, &segmentationError.ServerInternalError{Err: result_t.Error}
	}
	return result, nil
}

func (r *resultPosgresRepository) InsertResultData(in *entities.Result) error {
	result_t := r.db.Create(in)

	if result_t.Error != nil {
		log.Errorf("InsertDentistData:%v", result_t.Error)
		return result_t.Error
	}

	log.Debugf("InsertCockroachData: %v", result_t.RowsAffected)
	return nil
}

func (r *resultPosgresRepository) UpdateResultData(in *entities.Result) error {
	result_t := r.db.Save(in)

	if result_t.Error != nil {
		log.Errorf("UpdateDentistData:%v", result_t.Error)
		return result_t.Error
	}

	log.Debugf("UpdateCockroachData: %v", result_t.RowsAffected)
	return nil
}

func (r *resultPosgresRepository) DeleteResultData(id *uint64) error {
	result_t := r.db.Model(&entities.Result{}).Where("id = ?", *id).Where("status <> ?", "Removed").Update("status", "Removed")

	if result_t.Error != nil {
		log.Errorf("DeleteDentistData:%v", result_t.Error)
		return result_t.Error
	}

	log.Debugf("DeleteCockroachData: %v", result_t.RowsAffected)
	return nil
}
