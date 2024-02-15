package repositories

import "segmentation/dentists/entities"

type DentistRepository interface {
	SearchUsername(username *string) (bool, error)
	InsertDentistData(in *entities.InsertDentist) error
}
