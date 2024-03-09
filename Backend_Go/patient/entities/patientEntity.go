package entities

import (
	dentist "segmentation/dentists/entities"

	"gorm.io/gorm"
)

type (
	Patient struct {
		gorm.Model
		DentistID uint
		Dentist   dentist.Dentist
		BirthDate string  `json:"birth_date"`
		Gender    string  `json:"gender"`
		Age       uint64  `json:"age"`
		Phone     *string `json:"phone"`
		Status    string  `json:"status"`
	}
)
