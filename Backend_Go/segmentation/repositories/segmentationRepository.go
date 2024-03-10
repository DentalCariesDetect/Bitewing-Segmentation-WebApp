package repositories

import "segmentation/segmentation/entities"

type ImageRepository interface {
	Search(key string, value *string) (bool, error)
	GetImageDataByKey(key string, value *string) (*entities.Image, error)
	InsertImageData(in *entities.Image) error
	UpdateImageData(in *entities.Image) error
	DeleteImageData(id *uint64) error
}

type BitewingRepository interface {
	Search(key string, value *string) (bool, error)
	GetBitewingDataByKey(key string, value *string) (*entities.Bitewing, error)
	InsertBitewingData(in *entities.Bitewing) error
	UpdateBitewingData(in *entities.Bitewing) error
	DeleteBitewingData(id *uint64) error
}

type ResultRepository interface {
	Search(key string, value *string) (bool, error)
	GetResultDataByKey(key string, value *string) (*entities.Result, error)
	InsertResultData(in *entities.Result) error
	UpdateResultData(in *entities.Result) error
	DeleteResultData(id *uint64) error
}

type ToothRepository interface {
	Search(key string, value *string) (bool, error)
	GetToothDataByKey(key string, value *string) (*entities.Tooth, error)
	InsertToothData(in *entities.Tooth) error
	UpdateToothData(in *entities.Tooth) error
	DeleteToothData(id *uint64) error
}
