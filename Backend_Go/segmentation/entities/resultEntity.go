package entities

import (
	"gorm.io/gorm"
)

type (
	Result struct {
		gorm.Model
		BitewingId uint32 `json:"bitewing_id" gorm:"foreignkey:Bitewing"`
		DentistId  uint32 `json:"dentist_id" gorm:"foreignkey:Dentist"`
		PatientId  uint32 `json:"patient_id" gorm:"foreignkey:Patient"`
		Status     string `json:"status"`
	}
)
