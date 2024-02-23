package entities

import (
	"gorm.io/gorm"
)

type (
	Tooth struct {
		gorm.Model
		ResultId   uint32  `json:"result_id" gorm:"foreignkey:Result"`
		ImageId    uint32  `json:"image_id" gorm:"foreignkey:Image"`
		Name       string  `json:"name"`
		TypeTooth  *string `json:"type_tooth"`
		TypeCaries *string `json:"type_caries"`
		Numbering  uint32  `json:"numbering"`
		Filing     *string `json:"filing"`
		Position   *string `json:"position"`
		Status     string  `json:"status"`
	}
)
