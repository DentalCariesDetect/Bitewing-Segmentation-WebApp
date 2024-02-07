package repositories

import "segmentation/dentists/entities"

type DentistRepository interface {
	InsertDentistData(in *entities.InsertDentist) error
}
