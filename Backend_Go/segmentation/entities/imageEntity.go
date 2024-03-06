package entities

import (
	"gorm.io/gorm"
)

type (
	Image struct {
		gorm.Model
		Size_x uint32 `json:"size_x"`
		Size_y uint32 `json:"size_y"`
		Path   string `json:"path"`
		Status string `json:"status"`
	}
)
