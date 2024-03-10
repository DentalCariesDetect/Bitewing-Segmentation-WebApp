package entities

import (
	dentist "segmentation/dentists/entities"
	patient "segmentation/patient/entities"

	"gorm.io/gorm"
)

type (
	Result struct {
		gorm.Model
		BitewingId  uint64
		Bitewing    Bitewing
		DentistId   uint64
		Dentist     dentist.Dentist
		PatientId   uint64
		Patient     patient.Patient
		Description string `json:"description"`
		Treatment   string `json:"treatment"`
		Status      string `json:"status"`
	}
)
