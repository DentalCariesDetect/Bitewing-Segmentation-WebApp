package entities

import (
	"gorm.io/gorm"
)

type (
	Tooth struct {
		gorm.Model
		ResultId   uint32
		Result     Result
		ImageId    uint32
		Image      Image
		Name       string  `json:"name"`
		TypeTooth  *string `json:"type_tooth"`
		TypeCaries *string `json:"type_caries"`
		Numbering  uint32
		Filing     *string `json:"filing"`
		Position   *string `json:"position"`
		Status     string  `json:"status"`
	}
)
