package repositories

import (
	"segmentation/segmentation/entities"
	segmentationError "segmentation/segmentation/errors"
	"time"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type imagePosgresRepository struct {
	db *gorm.DB
}

func NewImagePostgresRepository(db *gorm.DB) ImageRepository {
	return &imagePosgresRepository{db: db}
}

func (r *imagePosgresRepository) Search(key string, value *string) (bool, error) {
	image := new(entities.Image)
	result := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").Limit(1).Find(image)
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

func (r *imagePosgresRepository) GetImageDataByKey(key string, value *string) (*entities.Image, error) {
	image := new(entities.Image)
	dentist_data := r.db.Where(key+"= ?", *value).Where("status <> ?", "Removed").First(image)
	if dentist_data.Error != nil {
		return nil, &segmentationError.ServerInternalError{Err: dentist_data.Error}
	}
	return image, nil
}

func (r *imagePosgresRepository) InsertImageData(in *entities.Image) error {
	result := r.db.Create(in)

	if result.Error != nil {
		log.Errorf("InsertDentistData:%v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *imagePosgresRepository) UpdateImageData(in *entities.Image) error {
	result := r.db.Save(in)

	if result.Error != nil {
		log.Errorf("UpdateDentistData:%v", result.Error)
		return result.Error
	}

	log.Debugf("UpdateCockroachData: %v", result.RowsAffected)
	return nil
}

func (r *imagePosgresRepository) DeleteImageData(id *uint64) error {
	result := r.db.Model(&entities.Image{}).Where("id = ?", *id).Where("status <> ?", "Removed").Updates(map[string]interface{}{
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
