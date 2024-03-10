package entities

import (
	"gorm.io/gorm"
)

type (
	Patient struct {
		gorm.Model
		BirthDate string  `json:"birth_date"`
		Gender    string  `json:"gender"`
		Age       uint32  `json:"age"`
		Phone     *string `json:"phone"`
		Status    string  `json:"status"`
	}
)
